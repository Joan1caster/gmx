package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PlatformAsset struct {
	gorm.Model
	Currency      string          `gorm:"type:varchar(20)"`   // 货币类型
	TotalBalance  decimal.Decimal `gorm:"type:decimal(32,8)"` // 总余额
	InsuranceFund decimal.Decimal `gorm:"type:decimal(32,8)"` // 保险基金
	FeeRevenue    decimal.Decimal `gorm:"type:decimal(32,8)"` // 手续费收入
	HotWallet     decimal.Decimal `gorm:"type:decimal(32,8)"` // 热钱包
}

type UserAsset struct {
	gorm.Model
	UserID       uint            `gorm:"index"`              // 用户ID
	Currency     string          `gorm:"type:varchar(20)"`   // 货币类型
	Available    decimal.Decimal `gorm:"type:decimal(32,8)"` // 可用余额
	Frozen       decimal.Decimal `gorm:"type:decimal(32,8)"` // 冻结金额
	TotalBalance decimal.Decimal `gorm:"type:decimal(32,8)"` // 总余额
}
