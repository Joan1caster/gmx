package positionkeeper

import (
	"gmxBackend/config"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadConfig() {
	err := config.LoadConfig("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}

func GetClient() *ethclient.Client {
	client, err := ethclient.Dial(config.AppConfig.Account.NodeAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return client
}

// func getPositionQueueLengths() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
// 	cilent := GetClient()
// 	positionRouterCaller, _ := positionrouter.NewPositionRouterCaller(common.HexToAddress(config.AppConfig.Contract.PositionRouter), cilent)
// 	return positionRouterCaller.GetRequestQueueLengths(&bind.CallOpts{})
// }

// func GetIncreasePosition() {
// 	client := GetClient()
// 	privateKey, err := crypto.HexToECDSA(config.AppConfig.Account.PrivateKey)

// 	positionRouterCaller, _ := positionrouter.NewPositionRouterCaller(common.HexToAddress(config.AppConfig.Contract.PositionRouter), cilent)

// 	positionRouterCaller.GetIncreasePositionRequest
// }
