package blockChain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gmxBackend/config"
	"gmxBackend/contracts/core/orderbook"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client *ethclient.Client
)

type Account struct {
	accoutAdreass common.Address
}

func GetClient() (*ethclient.Client, error) {
	var err error
	client, err = ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ExecuteIncreaseOrder(useraddress common.Address, orderIndex *big.Int) (*types.Transaction, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}

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

func ExecuteDecreaseOrder(useraddress common.Address, orderIndex *big.Int) (*types.Transaction, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}

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
	tx, err := orderbooksessioin.ExecuteDecreaseOrder(useraddress, orderIndex, common.HexToAddress(config.AppConfig.Contract.GOVAddress))
	return tx, err
}

func GetIncreaseOrderT(userAddress common.Address, orderIndex *big.Int) {
	client, _ := GetClient()
	caller, _ := orderbook.NewOrderBookCaller(common.HexToAddress(config.AppConfig.Contract.OrderBook), client)
	fmt.Println(caller.GetIncreaseOrder(&bind.CallOpts{}, userAddress, orderIndex))
}
