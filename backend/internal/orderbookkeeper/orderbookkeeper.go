package orderbookkeeper

import (
	"fmt"
	orderbook "gmxBackend/contracts/core/orderbook"
)

func init() {
	fmt.Println("OrderBookKeeper package initialized")
}

func NewOrderBookKeeper() *orderbook.OrderBook {
	return &orderbook.OrderBook{}
}

func Subscribe() {
	fmt.Println("OrderBookKeeper Subscribe called")
}

func (o *orderbook.OrderBook) GetOrderBook() {
	fmt.Println("OrderBookKeeper GetOrderBook called")x
}

func (o *orderbook.OrderBook) AddOrder() {
	fmt.Println("OrderBookKeeper AddOrder called")
}


package main

import (
	"context"
	"gmxBack/abi/core/orderbook"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)



func main() {
	var (
		ctx         = context.Background()
		url         = "http://127.0.0.1:8545/"
		client, err = ethclient.DialContext(ctx, url)
	)
	filter, err := orderbook.NewOrderBookFilterer(common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"), client)
	if err != nil {
		log.Fatal(err)
	}
	var 
	filter.WatchCreateIncreaseOrder(
		&bind.WatchOpts{},
	)

}
