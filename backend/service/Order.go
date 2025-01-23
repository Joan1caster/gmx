package service

import (
	"context"
	"fmt"
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

// 处理价格更新
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
			_, err := orderbookkeeper.ExecuteOrder(common.HexToAddress(order.Account), orderIndex)
			if err != nil {
				return err
			}
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

// 处理订单事件
func HandlerOrderInfo(db *gorm.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 建立基本连接
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		return fmt.Errorf("connect to block chain failed: %w", err)
	}

	log.Println(config.GetString("OrderBook"))
	filter, err := orderbook.NewOrderBookFilterer(common.HexToAddress(config.GetString("OrderBook")), client)
	if err != nil {
		return fmt.Errorf("create filter failed: %w", err)
	}

	// 初始化管道
	account := []common.Address{}
	var x uint64 = 0
	increaseSink := make(chan *orderbook.OrderBookCreateIncreaseOrder)
	decreaseSink := make(chan *orderbook.OrderBookCreateDecreaseOrder)
	exeIncreaseSink := make(chan *orderbook.OrderBookExecuteIncreaseOrder)
	exeDecreaseSink := make(chan *orderbook.OrderBookExecuteDecreaseOrder)

	// 订阅所有事件
	increaseSub, err := filter.WatchCreateIncreaseOrder(&bind.WatchOpts{Start: &x}, increaseSink, account)
	if err != nil {
		return fmt.Errorf("WatchCreateIncreaseOrder failed: %w", err)
	}
	defer increaseSub.Unsubscribe()

	decreaseSub, err := filter.WatchCreateDecreaseOrder(&bind.WatchOpts{Start: &x}, decreaseSink, account)
	if err != nil {
		return fmt.Errorf("WatchCreateDecreaseOrder failed: %w", err)
	}
	defer decreaseSub.Unsubscribe()

	exeIncreaseSub, err := filter.WatchExecuteIncreaseOrder(&bind.WatchOpts{Start: &x}, exeIncreaseSink, account)
	if err != nil {
		return fmt.Errorf("WatchExecuteIncreaseOrder failed: %w", err)
	}
	defer exeIncreaseSub.Unsubscribe()

	exeDecreaseSub, err := filter.WatchExecuteDecreaseOrder(&bind.WatchOpts{Start: &x}, exeDecreaseSink, account)
	if err != nil {
		return fmt.Errorf("WatchExecuteDecreaseOrder failed: %w", err)
	}
	defer exeDecreaseSub.Unsubscribe()

	log.Println("Start to handle order info")

	// 处理创建增仓订单
	go func() {
		for event := range increaseSink {
			log.Println("Received create increase order:", event)
			repository.NewOrderRepository(db).CreateIncreaseOrder(*event)
		}
	}()

	// 处理创建减仓订单
	go func() {
		for event := range decreaseSink {
			log.Println("Received create decrease order:", event)
			repository.NewOrderRepository(db).CreateDecreaseOrder(*event)
		}
	}()

	// 处理执行增仓订单
	go func() {
		for event := range exeIncreaseSink {
			log.Println("Received execute increase order:", event)
			repository.NewOrderRepository(db).DeleteOrder(*event)
		}
	}()

	// 处理执行减仓订单
	go func() {
		for event := range exeDecreaseSink {
			log.Println("Received execute decrease order:", event)
			repository.NewOrderRepository(db).DeleteOrder(*event)
		}
	}()

	<-ctx.Done()
	return nil
}
