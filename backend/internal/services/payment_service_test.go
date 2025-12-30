package services

import (
	"errors"
	"esapp/internal/dto"
	"testing"
)

func TestGetPaymentHistory_Success(t *testing.T) {
	mockHistory := []dto.PaymentHistoryRespone{
		{
			GroupName: "Goa Trip",
			FromUser:  "Rahul",
			FromEmail: "rahul@test.com",
			Amount:    500,
			Direction: "paid",
		},
		{
			GroupName: "Office Lunch",
			FromUser:  "Anita",
			FromEmail: "anita@test.com",
			Amount:    300,
			Direction: "received",
		},
	}

	repo := &mockSettlementRepo{
		paymentHis:     mockHistory,
		payementHisErr: nil,
	}

	service := NewPaymentService(repo)

	result, err := service.GetPaymentHistory(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("expected 2 records, got %d", len(result))
	}

	if result[0].GroupName != "Goa Trip" {
		t.Fatalf("unexpected group name %s", result[0].GroupName)
	}
}

func TestGetPaymentHistory_Error(t *testing.T) {
	repo := &mockSettlementRepo{
		paymentHis:     nil,
		payementHisErr: errors.New("db error"),
	}

	service := NewPaymentService(repo)

	_, err := service.GetPaymentHistory(1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
