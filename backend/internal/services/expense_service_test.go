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

	splits []models.ExpenseSplitWithExpense
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

func (m *mockExpenseRepo) GetSplitsByGroupID(groupID uint) ([]models.ExpenseSplitWithExpense, error) {
	return m.splits, nil
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

func TestCalculateBalancesSuccess(t *testing.T) {
	expenseRepo := &mockExpenseRepo{
		splits: []models.ExpenseSplitWithExpense{
			{UserID: 1, Amount: 500, PaidByID: 1},
			{UserID: 2, Amount: 500, PaidByID: 1},
			{UserID: 1, Amount: 200, PaidByID: 2},
			{UserID: 2, Amount: 200, PaidByID: 2},
		},
	}

	groupRepo := &mockGroupRepo{
		requesterIsMember: true,
	}

	service := NewExpenseService(expenseRepo, groupRepo)

	balances, err := service.CalculateBalances(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if balances[1] != 300 {
		t.Fatalf("expected user 1 balance 300, got %v", balances[1])
	}

	if balances[2] != -300 {
		t.Fatalf("expected user 2 balance -300, got %v", balances[2])
	}

}

func TestCreateExpense_CustomSplitSuccess(t *testing.T) {
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
		Splits: []dto.SplitInput{
			{UserID: 1, Amount: 300},
			{UserID: 2, Amount: 700},
		},
	}

	err := service.CreateExpense(1, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCreateExpense_CustomSplitInvalidSum(t *testing.T) {
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
		Splits: []dto.SplitInput{
			{UserID: 1, Amount: 500},
			{UserID: 2, Amount: 400},
		},
	}

	err := service.CreateExpense(1, req)
	if err == nil {
		t.Fatal("expected error for invalid split sum")
	}
}
