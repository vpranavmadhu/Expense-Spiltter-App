package services

import (
	"esapp/internal/dto"
	"esapp/models"
	"testing"
)

type mockExpenseRepo struct {
	createErr error
	splitErr  error
}

func (m *mockExpenseRepo) CreateExpense(expense *models.Expense) error {
	expense.ID = 1
	return m.createErr
}

func (m *mockExpenseRepo) CreateSplits(splits []models.ExpenseSplit) error {
	return m.splitErr
}

func TestCreateExpense_Success(t *testing.T) {

	expenseRepo := &mockExpenseRepo{}
	groupRepo := &mockGroupRepo{
		requesterIsMember: true,
		members: []models.User{
			{ID: 1},
			{ID: 2},
		},
	}

	service := NewExpenseService(expenseRepo, groupRepo)

	req := dto.CreateExpenseRequest{
		GroupID: 10,
		Title:   "Dinner",
		Amount:  1000,
	}

	err := service.CreateExpense(1, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}
