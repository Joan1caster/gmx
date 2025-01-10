package orderbook

import (
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestMain(t *testing.T) {
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatal("dial failed")
	}
	filter, _ := NewOrderBookFilterer(common.HexToAddress("0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB"), client)
	sink := make(chan *OrderBookCreateIncreaseOrder)
	account := []common.Address{}
	filter.WatchCreateIncreaseOrder(&bind.WatchOpts{}, sink, account)
	go func() {
		for event := range sink {
			fmt.Printf("Received event: %+v\n", event)
		}
	}()

	it, _ := filter.FilterCreateIncreaseOrder(&bind.FilterOpts{}, account)
	fmt.Print(it.Next())
}
