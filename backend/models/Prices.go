package models

import (
	"sync/atomic"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Price struct {
	gorm.Model
	Symbol    string          `gorm:"type:varchar(20);index"` // 交易对
	Price     decimal.Decimal `gorm:"type:decimal(32,8)"`     // 当前价格
	Timestamp int64           `gorm:"type:bigint"`            // 时间戳
	Source    string          `gorm:"type:varchar(50)"`       // 价格来源
	Volume24h decimal.Decimal `gorm:"type:decimal(32,8)"`     // 24小时成交量
}

type Buffer struct {
	Prices map[string]Price // token地址 -> 价格数据的映射
}

type PriceBuffer struct {
	Buffers     [2]*Buffer   // 两个缓冲区
	ActiveIndex atomic.Value // 当前活跃缓冲区的索引
	UpdateCount uint64       // 更新计数器
}

func NewPriceBuffer() *PriceBuffer {
	pb := &PriceBuffer{
		Buffers: [2]*Buffer{
			{Prices: make(map[string]Price)},
			{Prices: make(map[string]Price)},
		},
	}
	pb.ActiveIndex.Store(0)
	return pb
}
