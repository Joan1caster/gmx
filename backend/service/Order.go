package service

import (
	"fmt"
	"gmxBackend/contracts/core/orderbook"
	"gmxBackend/internal/orderbookkeeper"
	rabbitmq "gmxBackend/middleware/mq"
	"gmxBackend/repository"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func HandlerPriceInfo(msg []byte) error {
	Price_Order := rabbitmq.Consumers["PriceUpdater"]

	// 设置消费者处理函数
	err := Price_Order.Consume(func(msg []byte) error {
		price := string(msg)

		for _, order := range repository.NewOrderRepository().GetIncreaseOrder(price) {
			orderbookkeeper.ExecuteOrder(order.Account, order.OrderIndex)
		}

		return nil
	}, "PriceOrder")
	if err != nil {
		log.Fatalf("Failed to start consuming on consumer1: %v", err)
	}
	return nil
}

func HandlerOrderInfo() {
	// 建立基本连接
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatal("dial failed")
	}
	filter, _ := orderbook.NewOrderBookFilterer(common.HexToAddress("0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB"), client)

	// 初始化管道
	account := []common.Address{}
	increaseSink := make(chan *orderbook.OrderBookCreateIncreaseOrder)
	decreaseSink := make(chan *orderbook.OrderBookCreateDecreaseOrder)

	increaseSub, err := filter.WatchCreateIncreaseOrder(&bind.WatchOpts{}, increaseSink, account)
	if err != nil {
		log.Fatal(err)
	}
	defer increaseSub.Unsubscribe()

	decreaseSub, err := filter.WatchCreateDecreaseOrder(&bind.WatchOpts{}, decreaseSink, account)
	if err != nil {
		log.Fatal(err)
	}
	defer decreaseSub.Unsubscribe()

	go func() {
		for {
			for event := range increaseSink {
				repository.NewOrderRepository().CreateIncreaseOrder(*event)
			}
		}
	}()

	go func() {
		for {
			for event := range decreaseSink {
				fmt.Printf("Received event: %+v\n", event)
			}
		}
	}()

}
