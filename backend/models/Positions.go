package models

import (
	"gorm.io/gorm"
)

// Order 表示订单事件的存储结构体
type Position struct {
	gorm.Model                   // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Account               string `gorm:"type:varchar(42);index"` // ETH 地址长度为42（包含0x前缀）
	OrderIndex            string `gorm:"type:varchar(78)"`       // 存储big.Int的字符串表示
	PurchaseToken         string `gorm:"type:varchar(42)"`
	PurchaseTokenAmount   string `gorm:"type:varchar(78)"`
	CollateralToken       string `gorm:"type:varchar(42)"`
	IndexToken            string `gorm:"type:varchar(42)"`
	SizeDelta             string `gorm:"type:varchar(78)"`
	IsLong                bool
	TriggerPrice          string `gorm:"type:varchar(78)"`
	TriggerAboveThreshold bool
	ExecutionFee          string `gorm:"type:varchar(78)"`
	ExecutionPrice        string `gorm:"type:varchar(78)"`
	BlockNumber           uint64 `gorm:"index"`
	BlockHash             string `gorm:"type:varchar(66)"`
	TxHash                string `gorm:"type:varchar(66);index"`
	TxIndex               uint
}
