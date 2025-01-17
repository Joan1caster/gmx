package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	UserID        uint            `gorm:"index"`              // 用户ID
	Symbol        string          `gorm:"type:varchar(20)"`   // 交易对
	Side          string          `gorm:"type:varchar(10)"`   // 方向：LONG/SHORT
	Size          decimal.Decimal `gorm:"type:decimal(32,8)"` // 持仓数量
	EntryPrice    decimal.Decimal `gorm:"type:decimal(32,8)"` // 开仓均价
	Leverage      int             `gorm:"type:int"`           // 杠杆倍数
	MarginUsed    decimal.Decimal `gorm:"type:decimal(32,8)"` // 占用保证金
	UnrealizedPnL decimal.Decimal `gorm:"type:decimal(32,8)"` // 未实现盈亏
	LiqPrice      decimal.Decimal `gorm:"type:decimal(32,8)"` // 强平价格
	Status        string          `gorm:"type:varchar(20)"`   // 状态：OPEN/CLOSED
}
