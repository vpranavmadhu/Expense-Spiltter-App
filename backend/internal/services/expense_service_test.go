package services

import (
	"esapp/internal/dto"
	"esapp/models"
	"testing"
)

type mockExpenseRepo struct {
	createErr error
	splitErr  error

	expenses    []models.Expense
	expensesErr error
}

func (m *mockExpenseRepo) CreateExpense(expense *models.Expense) error {
	expense.ID = 1
	return m.createErr
}

func (m *mockExpenseRepo) CreateSplits(splits []models.ExpenseSplit) error {
	return m.splitErr
}

func (m *mockExpenseRepo) GetExpensesByGroupID(groupID uint) ([]models.Expense, error) {
	return m.expenses, m.createErr
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

func TestListExpensesSuccess(t *testing.T) {

	expenseRepo := &mockExpenseRepo{
		expenses: []models.Expense{
			{ID: 1, Title: "Dinner", Amount: 1000, PaidByID: 1},
			{ID: 2, Title: "Taxi", Amount: 300, PaidByID: 2},
		},
	}

	groupRepo := &mockGroupRepo{
		requesterIsMember: true,
	}

	service := NewExpenseService(expenseRepo, groupRepo)

	expenses, err := service.ListExpenses(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(expenses) != 2 {
		t.Fatalf("expected 2 expenses, got %d", len(expenses))
	}
}
