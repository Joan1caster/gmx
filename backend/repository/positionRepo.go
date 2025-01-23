package repository

import (
	"gmxBackend/contracts/core/orderbook"
	"gmxBackend/models"

	"gorm.io/gorm"
)

type PositionRepository struct {
	db *gorm.DB
}

func NewPositionReposttory(db_ *gorm.DB) *PositionRepository {
	return &PositionRepository{db: db_}
}

func (r *PositionRepository) CreatePosition(positionEvent *orderbook.OrderBookExecuteIncreaseOrder) error {
	position := &models.Position{
		Account:               positionEvent.Account.Hex(),
		OrderIndex:            positionEvent.OrderIndex.String(),
		PurchaseToken:         positionEvent.PurchaseToken.Hex(),
		PurchaseTokenAmount:   positionEvent.PurchaseTokenAmount.String(),
		CollateralToken:       positionEvent.CollateralToken.Hex(),
		IndexToken:            positionEvent.IndexToken.Hex(),
		SizeDelta:             positionEvent.SizeDelta.String(),
		IsLong:                positionEvent.IsLong,
		TriggerPrice:          positionEvent.TriggerPrice.String(),
		TriggerAboveThreshold: positionEvent.TriggerAboveThreshold,
		ExecutionFee:          positionEvent.ExecutionFee.String(),
		ExecutionPrice:        positionEvent.ExecutionPrice.String(),
		BlockNumber:           positionEvent.Raw.BlockNumber,
		BlockHash:             positionEvent.Raw.BlockHash.Hex(),
		TxHash:                positionEvent.Raw.TxHash.Hex(),
		TxIndex:               positionEvent.Raw.TxIndex,
	}

	return r.db.Create(position).Error
}

func (r *PositionRepository) GetByID(id uint) (*models.Position, error) {
	var position models.Position
	err := r.db.First(&position, id).Error
	if err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *PositionRepository) GetByUserIDAndSymbol(userID uint, symbol string) (*models.Position, error) {
	var position models.Position
	err := r.db.Where("user_id = ? AND symbol = ? AND status = ?", userID, symbol, "OPEN").First(&position).Error
	if err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *PositionRepository) ListByUserID(userID uint) ([]models.Position, error) {
	var positions []models.Position
	err := r.db.Where("user_id = ?", userID).Find(&positions).Error
	if err != nil {
		return nil, err
	}
	return positions, nil
}

// 更新多个字段
func (r *PositionRepository) UpdatePosition(position *orderbook.OrderBookExecuteDecreaseOrder) error {
	return r.db.Save(position).Error
}

// 更新特定字段
func (r *PositionRepository) UpdateFields(position *models.Position, updates map[string]interface{}) error {
	return r.db.Model(position).Updates(updates).Error
}

// 删除指定信息
func (r *PositionRepository) Delete(position *models.Position) error {
	return r.db.Delete(position).Error
}

// 获取所有开仓的头寸
func (r *PositionRepository) ListOpenPositions() ([]models.Position, error) {
	var positions []models.Position
	err := r.db.Find(&positions).Error
	if err != nil {
		return nil, err
	}
	return positions, nil
}
