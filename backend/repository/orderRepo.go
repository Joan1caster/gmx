package repository

import (
	"gmxBackend/contracts/core/orderbook"
	"gmxBackend/models"

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
		Account:               order.Account,
		OrderIndex:            order.OrderIndex,
		PurchaseToken:         order.PurchaseToken,
		PurchaseTokenAmount:   order.PurchaseTokenAmount,
		CollateralToken:       order.CollateralToken,
		IndexToken:            order.IndexToken,
		SizeDelta:             order.SizeDelta,
		IsLong:                order.IsLong,
		TriggerPrice:          order.TriggerPrice,
		TriggerAboveThreshold: order.TriggerAboveThreshold,
		ExecutionFee:          order.ExecutionFee,
	}
	result := db.db.Create(&orderModel)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *OrderRepository) GetIncreaseOrderByPrice(price int64) ([]models.Order, error) {
	var orders []models.Order
	result := db.db.Where("trigger_price <= ? and islong = ?", price, true).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (db *OrderRepository) GetDecreaseOrderByPrice(price int64) ([]models.Order, error) {
	var orders []models.Order
	result := db.db.Where("trigger_price >= ? and islong = ?", price, false).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (db *OrderRepository) CreateDecreaseOrder(order orderbook.OrderBookCreateDecreaseOrder) error {
	orderModel := &models.Order{
		Account:               order.Account,
		OrderIndex:            order.OrderIndex,
		CollateralToken:       order.CollateralToken,
		CollateralDelta:       order.CollateralDelta,
		IndexToken:            order.IndexToken,
		SizeDelta:             order.SizeDelta,
		IsLong:                order.IsLong,
		TriggerPrice:          order.TriggerPrice,
		TriggerAboveThreshold: order.TriggerAboveThreshold,
		ExecutionFee:          order.ExecutionFee,
	}
	result := db.db.Create(&orderModel)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
