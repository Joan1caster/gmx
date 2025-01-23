package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Account               common.Address `gorm:"type:varchar(32)""` // 用户ID
	OrderIndex            *big.Int       `gorm:"type:int"`
	PurchaseToken         common.Address `gorm:"type:varchar(20)"`
	PurchaseTokenAmount   *big.Int       `gorm:"type:int"`
	CollateralToken       common.Address `gorm:"type:varchar(20)"`
	CollateralDelta       *big.Int       `gorm:"type:int"`
	IndexToken            common.Address `gorm:"type:varchar(20)"`
	SizeDelta             *big.Int       `gorm:"type:int"`
	IsLong                bool           `gorm:"type:bool"`
	TriggerPrice          *big.Int       `gorm:"type:int"`
	TriggerAboveThreshold bool           `gorm:"type:bool"`
	ExecutionFee          *big.Int       `gorm:"type:int"`
}
