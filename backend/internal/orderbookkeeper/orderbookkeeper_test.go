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
	SubscriptionEvent(common.HexToAddress("0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB"))
	a := orderbook.OrderBookABI
	fmt.Print(a)
}
