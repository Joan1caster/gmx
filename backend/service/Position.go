package service

import (
	"errors"
	"fmt"
	"gmxBackend/blockChain"
	"gmxBackend/config"
	"gmxBackend/contracts/core/orderbook"
	rabbitmq "gmxBackend/middleware/mq"
	"gmxBackend/repository"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type PositionService struct {
	positionRepo   *repository.PositionRepository
	orderRepo      *repository.OrderRepository
	priceKeeper    *blockChain.PriceKeeper
	positionKeeper *blockChain.PositionKeeper
}

func NewPositionService(orRe *repository.OrderRepository, poRe *repository.PositionRepository) *PositionService {
	return &PositionService{positionRepo: poRe, orderRepo: orRe}
}

// 订单执行头寸入库
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
			p.positionRepo.CreatePosition(event)
		}
	}()

	go func() {
		platform := rabbitmq.Consumers["PricePlatform"]
		platform.Consume(p.CheckClosing, "PriceUser")
	}()
	return nil
}

func (p *PositionService) CheckClosing(TokenPrice []byte) error {
	// 检查爆仓
	Postions, err := p.positionRepo.ListOpenPositions()
	if err != nil {
		return err
	}

	tokenPrice := new(big.Int)
	tokenPrice.SetBytes(TokenPrice)

	for _, Position := range Postions {
		TriggerPrice_ := new(big.Int)
		TriggerPrice_.SetString(Position.TriggerPrice, 30)
		size_ := new(big.Int)
		size_.SetString(Position.SizeDelta, 18)
		collateral_ := new(big.Int)
		collateral_.SetString(Position.PurchaseTokenAmount, 18)
		isLiquidation, err := calculateLiquidationCondition(collateral_, size_, tokenPrice, TriggerPrice_, Position.IsLong)
		if err != nil {
			return errors.New("calculate Liquidation Condition error:" + err.Error())
		}
		if isLiquidation {
			err = p.priceKeeper.UpdateBTCPrice(tokenPrice)
			if err != nil {
				return errors.New("Update blockchain BTCPrice error:" + err.Error())
			}
			p.positionKeeper.LiquidatePosition(
				common.HexToAddress(Position.Account),
				common.HexToAddress(Position.CollateralToken),
				common.HexToAddress(Position.IndexToken),
				Position.IsLong,
				common.HexToAddress(Position.Account),
			)
		}
	}

	return nil
}

func calculateLiquidationCondition(
	collateral_ *big.Int,
	size_ *big.Int,
	tokenPrice *big.Int,
	TriggerPrice *big.Int,
	IsLong bool,
) (bool, error) {
	//	collection < (triggerPrice - priceNow) * size + fundingFee + positionFee
	// positionFee = size/10000
	// fundingFee  = size * fundingrate / 1000000
	// fundingRate = vault.cumulativeFundingRates(_collateralToken).sub(_entryFundingRate); 使用估算值1000
	precision := new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil) // 10^30 精度
	divisor := big.NewInt(11000)

	// 1. 计算 size / 11000
	sizeDiv11000 := new(big.Int).Div(size_, divisor)

	// 2. 计算 price_now - price_old (priceDiff)
	priceDiff := new(big.Int)
	if IsLong {
		priceDiff = new(big.Int).Sub(tokenPrice, TriggerPrice)
	} else {
		priceDiff = new(big.Int).Sub(TriggerPrice, tokenPrice)
	}

	// 3. 计算 size * (price_now - price_old)
	// 注意：由于价格有30位精度，结果需要除以precision来调整到18位精度
	priceImpact := new(big.Int).Mul(size_, priceDiff)
	priceImpact = new(big.Int).Div(priceImpact, precision)

	// 4. 计算右边总和: (size / 11000 + size * (price_now - price_old))
	rightSide := new(big.Int).Add(sizeDiv11000, priceImpact)

	// 5. 比较 collateral < rightSide

	return collateral_.Cmp(rightSide) < 0, nil
}
