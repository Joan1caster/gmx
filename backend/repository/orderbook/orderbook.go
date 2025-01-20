package orderbook

import (
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func (db *OrderRepository) CreateOrder(order *Order) error {
	return db.db.Create(order).Error
}

func (db *OrderRepository) GetOrder(orderID int) (*Order, error) {
	order := &Order{}
	err := db.db.First(order, orderID).Error
	return order
}

func (db *OrderRepository) DelOrders() ([]Order, error) {
	orders := []Order{}
	err := db.db.Find(&orders).Error
	return orders, err
}
