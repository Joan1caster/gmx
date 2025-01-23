package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Account               string `gorm:"type:varchar(128)""`
	OrderIndex            string `gorm:"type:varchar(128)"`
	PurchaseToken         string `gorm:"type:varchar(128)"`
	PurchaseTokenAmount   string `gorm:"type:varchar(128)"`
	CollateralToken       string `gorm:"type:varchar(128)"`
	CollateralDelta       string `gorm:"type:varchar(128)"`
	IndexToken            string `gorm:"type:varchar(128)"`
	SizeDelta             string `gorm:"type:varchar(128)"`
	IsLong                bool   `gorm:"type:bool"`
	TriggerPrice          string `gorm:"type:string;index:idx_trigger_price,sort:desc"`
	TriggerAboveThreshold bool   `gorm:"type:bool;index"`
	ExecutionFee          string `gorm:"type:varchar(128)"`
	Type                  string `gorm:"type:varchar(20);check:type in ('increase','decrease')"`
}
