package orderbookkeeper

import (
	"context"
	"gmxBackend/config"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	// client := getClient()

	// privateKey, err := crypto.HexToECDSA(config.AppConfig.Account.PrivateKey)
	// if err != nil {
	// 	log.Fatalf("Failed to load private key: %v", err)
	// }

	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // ChainID for mainnet is 1
	// if err != nil {
	// 	log.Fatalf("Failed to create auth: %v", err)
	// }

	// contractAddress := common.HexToAddress(config.AppConfig.Contract.OrderBook)

	// exchange, err := bindings.NewSimpleExchange(contractAddress, client)
	// if err != nil {
	// 	log.Fatalf("Failed to instantiate the contract: %v", err)
	// }

	// // Example: Buy tokens
	// tx, err := exchange.BuyToken(auth)
	// if err != nil {
	// 	log.Fatalf("Failed to buy tokens: %v", err)
	// }

	// fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	// // Example: Get contract balance
	// balance, err := exchange.GetBalance(nil)
	// if err != nil {
	// 	log.Fatalf("Failed to get balance: %v", err)
	// }

	// fmt.Printf("Contract balance: %s\n", balance.String())
}

func SubscriptionEvent(contractAddress common.Address) {
	client := getClient()
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

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
			log.Printf("Log: %v", vLog)
		}
	}

}
