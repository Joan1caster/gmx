package service

import (
	"fmt"
	"gmxBackend/config"
	"gmxBackend/contracts/core/orderbook"
	"gmxBackend/repository"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type PositionService struct {
	positionRepo *repository.PositionRepository
	orderRepo    *repository.OrderRepository
}

func NewPositionService(orRe *repository.OrderRepository, poRe *repository.PositionRepository) *PositionService {
	return &PositionService{positionRepo: poRe, orderRepo: orRe}
}

func (p *PositionService) HandlerPositionInfo() error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	filter, err := orderbook.NewOrderBookFilterer(common.HexToAddress(config.GetString("OrderBook")), client)
	if err != nil {
		return fmt.Errorf("create filter failed: %w", err)
	}

	account := []common.Address{}
	var x uint64 = 0
	exeIncreaseSink := make(chan *orderbook.OrderBookExecuteIncreaseOrder)
	exeDecreaseSink := make(chan *orderbook.OrderBookExecuteDecreaseOrder)

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

	// 处理执行增仓订单
	go func() {
		for event := range exeIncreaseSink {
			log.Println("Received execute increase order:", event)
			p.orderRepo.DeleteOrder(event.Account, event.OrderIndex)

		}
	}()

	// 处理执行减仓订单
	go func() {
		for event := range exeDecreaseSink {
			log.Println("Received execute decrease order:", event)
			p.orderRepo.DeleteOrder(event.Account, event.OrderIndex)
		}
	}()
	return nil
}
