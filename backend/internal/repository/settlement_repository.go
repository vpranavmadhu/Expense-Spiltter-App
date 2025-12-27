package repository

import (
	"esapp/models"

	"gorm.io/gorm"
)

type SettlementRepository interface {
	Create(payment *models.SettlementPayment) error
	GetByGroupID(groupID uint) ([]models.SettlementPayment, error)
}

type settlementRepository struct {
	db *gorm.DB
}

func NewSettlementRepository(db *gorm.DB) SettlementRepository {
	return &settlementRepository{db: db}
}

func (r *settlementRepository) Create(p *models.SettlementPayment) error {
	return r.db.Create(p).Error
}

func (r *settlementRepository) GetByGroupID(groupID uint) ([]models.SettlementPayment, error) {
	var payments []models.SettlementPayment
	err := r.db.Where("group_id = ?", groupID).Find(&payments).Error
	return payments, err
}
