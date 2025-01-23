package service

import (
	"fmt"
	"gmxBackend/blockChain"
	"gmxBackend/config"
	"gmxBackend/contracts/core/orderbook"
	"gmxBackend/models"
	"gmxBackend/repository"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type PositionService struct {
	positionRepo *repository.PositionRepository
	orderRepo    *repository.OrderRepository
	positions    []models.Position
}

func NewPositionService(orRe *repository.OrderRepository, poRe *repository.PositionRepository) *PositionService {
	return &PositionService{positionRepo: poRe, orderRepo: orRe}
}

func (p *PositionService) HandlerPositionInfo() error {
	client, err := blockChain.GetClient()
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
			p.positionRepo.CreatePosition(event)
		}
	}()

	return nil
}
