package service

import (
	"fmt"
	"gmxBackend/repository"
	"math/big"
)

// 清算条件：
//
//	押金 < 亏损 + fundingFee + positionFee
//
// positionFee = size/10000
// fundingFee  = size * fundingrate / 1000000
// fundingRate = vault.cumulativeFundingRates(_collateralToken).sub(_entryFundingRate); 使用估算值1000
type RiskControl struct {
	positionRepo *repository.PositionRepository
}

func (r *RiskControl) CheckClosing(TokenPrice []byte) error {
	// 遍历所有仓位
	//		condition = collateral < (size / 11000 + size * (price_now - price_old))
	//		if condition
	//			vault.liquidatePosition
	//			DB: positions delete
	Postions, err := r.positionRepo.ListOpenPositions()
	if err != nil {
		return err
	}

	tokenPrice := new(big.Int)
	tokenPrice.SetBytes(TokenPrice)

	TriggerPrice := new(big.Int)

	for _, Position := range Postions {
		if Position.IsLong { // 做多
			TriggerPrice.SetString(Position.TriggerPrice, 30)
			priceDiff := new(big.Int).Sub(TriggerPrice, tokenPrice)
			fmt.Print(priceDiff)
		}
	}

	return nil
}
