package service

// 清算条件：
//	押金 < 亏损 + fundingFee + positionFee
// positionFee = size/10000
// fundingFee  = size * fundingrate / 1000000
// fundingRate = vault.cumulativeFundingRates(_collateralToken).sub(_entryFundingRate); 使用估算值1000
//
type RiskControl struct {
}

func (*RiskControl) CheckClosing(TokenPrice []byte) error {
	// 遍历所有仓位
	//		condition = collateral < (size / 11000 + size * (price_now - price_old))
	//		if condition
	//			vault.liquidatePosition
	//			DB: positions delete

	return nil
}
