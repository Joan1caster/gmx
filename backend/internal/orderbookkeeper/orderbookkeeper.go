package orderbookkeeper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gmxBackend/config"
	"gmxBackend/contracts/core/orderbook"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
	accoutAdreass common.Address
}

func LoadConfig() {
	err := config.LoadConfig("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}

func getClient() *ethclient.Client {
	client, err := ethclient.Dial(config.AppConfig.Account.NodeAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return client
}

func executeOrder() {
	client := getClient()

	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
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
	// orderbooksessioin.ExecuteIncreaseOrder(useraddress, order, govaddress)
}

func SubscriptionEvent(contractAddress common.Address) {
	client := getClient()
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	orderbookfilter, _ := orderbook.NewOrderBookFilterer(common.HexToAddress("0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB"), client)

	fmt.Println(*orderbookfilter)

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

func GetIncreaseOrderT() {
	client := getClient()
	caller, _ := orderbook.NewOrderBookCaller(common.HexToAddress(config.AppConfig.Contract.OrderBook), client)
	fmt.Println(caller.GetIncreaseOrder(&bind.CallOpts{}, common.HexToAddress("0x0"), big.NewInt(1)))
}
