package positionkeeper

import (
	"gmxBackend/config"
	positionrouter "gmxBackend/contracts/core/positionrouter"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

func getPositionQueueLengths() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	cilent := getClient()
	positionRouterCaller, _ := positionrouter.NewPositionRouterCaller(common.HexToAddress(config.AppConfig.Contract.PositionRouter), cilent)
	return positionRouterCaller.GetRequestQueueLengths(&bind.CallOpts{})
}

func updatePriceBitsAndOptionallyExecute() {

}
