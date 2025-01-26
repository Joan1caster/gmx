package repository

import (
	"gmxBackend/contracts/core/orderbook"
	"gmxBackend/models"
	"gmxBackend/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (db *OrderRepository) CreateIncreaseOrder(order orderbook.OrderBookCreateIncreaseOrder) error {
	orderModel := &models.Order{
		Account:               order.Account.Hex(),
		OrderIndex:            utils.Uint256ToString(order.OrderIndex, 18),
		PurchaseToken:         order.PurchaseToken.Hex(),
		PurchaseTokenAmount:   utils.Uint256ToString(order.PurchaseTokenAmount, 18),
		CollateralToken:       order.CollateralToken.String(),
		IndexToken:            order.IndexToken.Hex(),
		SizeDelta:             utils.Uint256ToString(order.SizeDelta, 18),
		IsLong:                order.IsLong,
		TriggerPrice:          utils.Uint256ToString(order.TriggerPrice, 30),
		TriggerAboveThreshold: order.TriggerAboveThreshold,
		ExecutionFee:          utils.Uint256ToString(order.ExecutionFee, 18),
		Type:                  "increase",
	}
	result := db.db.Create(&orderModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *OrderRepository) GetLessOrderByPrice(price float64) ([]models.Order, error) {
	var orders []models.Order
	result := db.db.Where("trigger_price < ? and trigger_above_threshold = ?", price, 1).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (db *OrderRepository) GetGreateOrderByPrice(price float64) ([]models.Order, error) {
	var orders []models.Order
	result := db.db.Where("trigger_price > ? and trigger_above_threshold = ?", price, 0).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (db *OrderRepository) CreateDecreaseOrder(order orderbook.OrderBookCreateDecreaseOrder) error {
	orderModel := &models.Order{
		Account:               order.Account.Hex(),
		OrderIndex:            utils.Uint256ToString(order.OrderIndex, 18),
		CollateralDelta:       utils.Uint256ToString(order.CollateralDelta, 18),
		CollateralToken:       order.CollateralToken.Hex(),
		IndexToken:            order.IndexToken.Hex(),
		SizeDelta:             utils.Uint256ToString(order.SizeDelta, 18),
		IsLong:                order.IsLong,
		TriggerPrice:          utils.Uint256ToString(order.TriggerPrice, 30),
		TriggerAboveThreshold: order.TriggerAboveThreshold,
		ExecutionFee:          utils.Uint256ToString(order.ExecutionFee, 18),
		Type:                  "decrease",
	}
	result := db.db.Create(&orderModel)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *OrderRepository) DeleteOrder(account common.Address, index *big.Int) error {
	var order models.Order
	result := db.db.Where("account = ? and order_index = ?",
		account.Hex(),
		utils.Uint256ToString(index, 18),
	).Find(&order)

	if result.Error != nil {
		return errors.New("Search Order Error:" + result.Error.Error())
	}

	return db.db.Delete(order).Error
}
