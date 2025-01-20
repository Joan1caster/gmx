package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func StringToUint256(str string, decimals int32) *big.Int {
	// 将字符串转换为 decimal
	d, _ := decimal.NewFromString(str)

	// 将数字乘以 10^decimals
	multiplier := decimal.NewFromInt(10).Pow(decimal.NewFromInt32(decimals))
	result := d.Mul(multiplier)

	// 转换为 big.Int
	bigInt := new(big.Int)
	bigInt.SetString(result.String(), 10)

	return bigInt
}

func Uint256ToString(num *big.Int, decimals int32) string {
	// 转换为 decimal
	d := decimal.NewFromBigInt(num, 0)

	// 除以 10^decimals
	divisor := decimal.NewFromInt(10).Pow(decimal.NewFromInt32(decimals))
	result := d.Div(divisor)

	return result.String()
}
