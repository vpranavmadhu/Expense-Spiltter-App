package repository

import (
	"esapp/internal/dto"
	"esapp/models"

	"gorm.io/gorm"
)

type SettlementRepository interface {
	Create(payment *models.SettlementPayment) error
	GetByGroupID(groupID uint) ([]models.SettlementPayment, error)
	GetPaymentHistoryByUser(userID uint) ([]dto.PaymentHistoryRespone, error)
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

func (r *settlementRepository) GetPaymentHistoryByUser(
	userID uint,
) ([]dto.PaymentHistoryRespone, error) {

	var results []dto.PaymentHistoryRespone

	err := r.db.Raw(`
		SELECT
			g.name AS group_name,

			CASE
				WHEN sp.from_user_id = ? THEN u_to.username
				ELSE u_from.username
			END AS from_user,

			CASE
				WHEN sp.from_user_id = ? THEN u_to.email
				ELSE u_from.email
			END AS from_email,

			sp.amount,
			sp.created_at,

			CASE
				WHEN sp.from_user_id = ? THEN 'paid'
				ELSE 'received'
			END AS direction

		FROM settlement_payments sp
		JOIN groups g ON g.id = sp.group_id
		JOIN users u_from ON u_from.id = sp.from_user_id
		JOIN users u_to ON u_to.id = sp.to_user_id
		WHERE sp.from_user_id = ? OR sp.to_user_id = ?
		ORDER BY sp.created_at DESC
	`, userID, userID, userID, userID, userID).
		Scan(&results).Error

	return results, err
}
