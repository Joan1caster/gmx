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
	SubscriptionEvent(common.HexToAddress("0x9E545E3C0baAB3E08CdfD552C960A1050f373042"))
	a := orderbook.OrderBookABI
	fmt.Print(a)
}
