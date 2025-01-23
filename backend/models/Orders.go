package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Account               string `gorm:"type:varchar(128)""` // 用户ID
	OrderIndex            string `gorm:"type:varchar(128)"`
	PurchaseToken         string `gorm:"type:varchar(128)"`
	PurchaseTokenAmount   string `gorm:"type:varchar(128)"`
	CollateralToken       string `gorm:"type:varchar(128)"`
	CollateralDelta       string `gorm:"type:varchar(128)"`
	IndexToken            string `gorm:"type:varchar(128)"`
	SizeDelta             string `gorm:"type:varchar(128)"`
	IsLong                bool   `gorm:"type:bool"`
	TriggerPrice          string `gorm:"type:varchar(128)"`
	TriggerAboveThreshold bool   `gorm:"type:bool"`
	ExecutionFee          string `gorm:"type:varchar(128)"`
}
