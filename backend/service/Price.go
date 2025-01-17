package service

import (
	"gmxBackend/models"
	"gmxBackend/repository"
	"sync/atomic"
)

type PriceService struct {
	priceBuffer *models.PriceBuffer
	priceRepo   *repository.PriceRepository
}

func NewPriceService(_priceBuffer *models.PriceBuffer, _priceRepo *repository.PriceRepository) *PriceService {
	return &PriceService{
		priceBuffer: _priceBuffer,
		priceRepo:   _priceRepo,
	}
}

func (p *PriceService) UpdatePrices(updates []models.Price) error {
	activeIdx := p.priceBuffer.ActiveIndex.Load().(int)
	// 计算非活跃缓冲区索引
	inactiveIdx := 1 - activeIdx

	// 获取非活跃缓冲区
	inactiveBuffer := p.priceBuffer.Buffers[inactiveIdx]

	// 复制当前活跃缓冲区的数据到非活跃缓冲区
	for k, v := range p.priceBuffer.Buffers[activeIdx].Prices {
		inactiveBuffer.Prices[k] = v
	}

	// 更新非活跃缓冲区的价格
	for _, update := range updates {
		inactiveBuffer.Prices[update.Symbol] = update
	}

	// 原子性地切换活跃缓冲区
	p.priceBuffer.ActiveIndex.Store(inactiveIdx)

	// 更新计数器
	atomic.AddUint64(&p.priceBuffer.UpdateCount, 1)

	p.priceRepo.UpdatePrice(updates)
	return nil
}
