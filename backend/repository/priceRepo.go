package repository

import "gmxBackend/models"

type PriceRepository struct {
}

func NewPriceRepo() *PriceRepository {
	return &PriceRepository{}
}

func (p *PriceRepository) UpdatePrice(prices []models.Price) error {
	return nil
}
