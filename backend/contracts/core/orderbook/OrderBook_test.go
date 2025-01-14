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
	var x uint64 = 0
	sub, err := filter.WatchCreateIncreaseOrder(&bind.WatchOpts{Start: &x}, sink, account)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		for event := range sink {
			fmt.Printf("Received event: %+v\n", event)
		}
	}

	order, err := NewOrderBookCaller(common.HexToAddress("0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB"), client)

	fmt.Println(order.IncreaseOrdersIndex(&bind.CallOpts{}, common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")))
	it, err := filter.FilterCreateIncreaseOrder(&bind.FilterOpts{Start: 0}, account)
	if err != nil {
		log.Fatal(err)
	}
	for range 25 {
		fmt.Println(it.Next())
		fmt.Printf("%+v\n", it.Event)
	}
	it.Close()
}
