package orderbookkeeper

import (
	"fmt"
	"gmxBackend/config"
	"gmxBackend/contracts/core/orderbook"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestMain(t *testing.T) {
	config.LoadConfig("../../config/config.yaml")
	SubscriptionEvent(common.HexToAddress(config.AppConfig.Contract.OrderBook))
	a := orderbook.OrderBookABI
	fmt.Print(a)
}
