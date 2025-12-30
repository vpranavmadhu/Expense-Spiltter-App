package services

import (
	"esapp/internal/dto"
	"esapp/internal/repository"
)

type PaymentService interface {
	GetPaymentHistory(userID uint) ([]dto.PaymentHistoryRespone, error)
}

type paymentService struct {
	settlementRepo repository.SettlementRepository
}

func NewPaymentService(settlementRepo repository.SettlementRepository) PaymentService {
	return &paymentService{settlementRepo: settlementRepo}
}

func (s *paymentService) GetPaymentHistory(userID uint) ([]dto.PaymentHistoryRespone, error) {

	return s.settlementRepo.GetPaymentHistoryByUser(userID)
}
