package orderbookkeeper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gmxBackend/config"
	"gmxBackend/contracts/core/orderbook"
	"gmxBackend/models"
	"gmxBackend/repository"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadConfig() {
	err := config.LoadConfig("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}

var (
	client *ethclient.Client
	once   sync.Once
)

type Account struct {
	accoutAdreass common.Address
}

func GetClient() *ethclient.Client {
	once.Do(func() {
		var err error
		client, err = ethclient.Dial(config.AppConfig.Account.NodeAddress)
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
	})
	return client
}

func ExecuteOrder(useraddress common.Address, orderIndex *big.Int) (*types.Transaction, error) {
	client := GetClient()

	privateKey, err := crypto.HexToECDSA(config.AppConfig.Account.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	orderbookins, err := orderbook.NewOrderBook(common.HexToAddress(config.AppConfig.Contract.OrderBook), client)
	if err != nil {
		log.Fatal(err)
	}
	orderbooksessioin := orderbook.OrderBookSession{
		Contract:     orderbookins,
		CallOpts:     bind.CallOpts{},
		TransactOpts: *auth,
	}
	fmt.Println(orderbooksessioin)
	tx, err := orderbooksessioin.ExecuteIncreaseOrder(useraddress, orderIndex, common.HexToAddress(config.AppConfig.Contract.GOVAddress))
	return tx, err
}

func SubscriptionEvent(contractAddress common.Address) {
	client := GetClient()
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{common.HexToHash("0xb27b9afe3043b93788c40cfc3cc73f5d928a2e40f3ba01820b246426de8fa1b9")}},
	}

	orderbookfilter, _ := orderbook.NewOrderBookFilterer(contractAddress, client)

	fmt.Println(*orderbookfilter)

	testabi, _ := orderbook.OrderBookMetaData.GetAbi()
	fmt.Println(testabi.Events)
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logs:
			// Handle the log
			log.Printf("Log: %+v", vLog)
			fmt.Println(vLog.Topics[0])
			event, _ := orderbookfilter.ParseCreateIncreaseOrder(vLog)
			log.Printf("LogEvent: %+v", event)
		}
	}
}

func GetIncreaseOrderT(userAddress common.Address, orderIndex *big.Int) {
	client := GetClient()
	caller, _ := orderbook.NewOrderBookCaller(common.HexToAddress(config.AppConfig.Contract.OrderBook), client)
	fmt.Println(caller.GetIncreaseOrder(&bind.CallOpts{}, userAddress, orderIndex))
}

func WaitTx(tx *types.Transaction, order models.Order, orderRepo *repository.OrderRepository) bool {
	client := GetClient()
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("等待交易确认失败:", err)
	}

	// 3. 检查交易状态
	if receipt.Status == types.ReceiptStatusSuccessful {
		return true
	} else {
		return false
	}
}
