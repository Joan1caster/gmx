package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      uint            `gorm:"index"`              // 用户ID
	OrderType   string          `gorm:"type:varchar(20)"`   // 订单类型：LIMIT/MARKET
	Side        string          `gorm:"type:varchar(10)"`   // 方向：LONG/SHORT
	Symbol      string          `gorm:"type:varchar(20)"`   // 交易对
	Size        decimal.Decimal `gorm:"type:decimal(32,8)"` // 数量
	Price       decimal.Decimal `gorm:"type:decimal(32,8)"` // 价格
	Leverage    int             `gorm:"type:int"`           // 杠杆倍数
	Status      string          `gorm:"type:varchar(20)"`   // 状态：PENDING/FILLED/CANCELED
	FilledSize  decimal.Decimal `gorm:"type:decimal(32,8)"` // 已成交数量
	FilledPrice decimal.Decimal `gorm:"type:decimal(32,8)"` // 成交价格
	Fee         decimal.Decimal `gorm:"type:decimal(32,8)"` // 手续费
}
