package service

import (
	"context"
	"gmxBackend/config"
	"gmxBackend/contracts/core/orderbook"
	"gmxBackend/internal/orderbookkeeper"
	rabbitmq "gmxBackend/middleware/mq"
	"gmxBackend/repository"
	"gmxBackend/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

func HandlerPriceInfo(db *gorm.DB) error {
	Price_Order := rabbitmq.Consumers["PriceUpdater"]

	// 设置消费者处理函数
	err := Price_Order.Consume(func(msg []byte) error {

		price := utils.StringToUint256(string(msg), 18)
		orders, err := repository.NewOrderRepository(db).GetLessOrderByPrice(price)
		if err != nil {
			log.Fatal(err)
		}

		for _, order := range orders {
			orderIndex := new(big.Int)
			orderIndex.SetString(order.OrderIndex, 10)
			orderbookkeeper.ExecuteOrder(common.HexToAddress(order.Account), orderIndex)
		}

		orders, err = repository.NewOrderRepository(db).GetLessOrderByPrice(price)
		if err != nil {
			log.Fatal(err)
		}

		for _, order := range orders {
			orderIndex := new(big.Int)
			orderIndex.SetString(order.OrderIndex, 10)
			orderbookkeeper.ExecuteOrder(common.HexToAddress(order.Account), orderIndex)
		}
		return nil
	}, "PriceOrder")
	if err != nil {
		log.Fatalf("Failed to start consuming on consumer1: %v", err)
	}
	return nil
}

func HandlerOrderInfo(db *gorm.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 建立基本连接
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatal("dial failed")
	}
	log.Println(config.AppConfig.Contract.OrderBook)
	filter, _ := orderbook.NewOrderBookFilterer(common.HexToAddress(config.AppConfig.Contract.OrderBook), client)

	// 初始化管道
	account := []common.Address{}
	var x uint64 = 0
	increaseSink := make(chan *orderbook.OrderBookCreateIncreaseOrder)
	decreaseSink := make(chan *orderbook.OrderBookCreateDecreaseOrder)

	increaseSub, err := filter.WatchCreateIncreaseOrder(&bind.WatchOpts{Start: &x}, increaseSink, account)
	if err != nil {
		log.Fatal(err)
	}
	defer increaseSub.Unsubscribe()

	decreaseSub, err := filter.WatchCreateDecreaseOrder(&bind.WatchOpts{Start: &x}, decreaseSink, account)
	if err != nil {
		log.Fatal(err)
	}
	defer decreaseSub.Unsubscribe()

	log.Println("Start to handle order info")

	log.Println("Start to handle increase order info")
	go func() {
		for {
			for event := range increaseSink {
				log.Println(event)
				repository.NewOrderRepository(db).CreateIncreaseOrder(*event)
			}
		}
	}()

	go func() {
		for {
			for event := range decreaseSink {
				log.Println(event)
				repository.NewOrderRepository(db).CreateDecreaseOrder(*event)
			}
		}
	}()
	<-ctx.Done()
	return nil
}
