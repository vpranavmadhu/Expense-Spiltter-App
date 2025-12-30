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

	expenseWithShare []dto.ExpenseResponse

	settleErr error
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

func (m *mockExpenseRepo) GetExpensesWithMyShare(groupID uint, userID uint) ([]dto.ExpenseResponse, error) {
	return m.expenseWithShare, nil
}

type mockSettlementRepo struct {
	createErr error
	payments  []models.SettlementPayment

	paymentHis     []dto.PaymentHistoryRespone
	payementHisErr error
}

func (m *mockSettlementRepo) Create(p *models.SettlementPayment) error {
	return m.createErr
}

func (m *mockSettlementRepo) GetByGroupID(groupID uint) ([]models.SettlementPayment, error) {
	return m.payments, nil
}

func (m *mockExpenseRepo) SettleExpenseSplit(expenseID, userID uint) error {
	return m.settleErr
}

func (m *mockSettlementRepo) GetPaymentHistoryByUser(userID uint) ([]dto.PaymentHistoryRespone, error) {
	return m.paymentHis, m.payementHisErr
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
	settlementRepo := &mockSettlementRepo{}

	service := NewExpenseService(expenseRepo, groupRepo, settlementRepo)

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

	settlementRepo := &mockSettlementRepo{}

	service := NewExpenseService(expenseRepo, groupRepo, settlementRepo)

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

	settlementRepo := &mockSettlementRepo{}

	service := NewExpenseService(expenseRepo, groupRepo, settlementRepo)

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

	settlementRepo := &mockSettlementRepo{}

	service := NewExpenseService(expenseRepo, groupRepo, settlementRepo)

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

	settlementRepo := &mockSettlementRepo{}

	service := NewExpenseService(expenseRepo, groupRepo, settlementRepo)

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

func TestCalculateBalances_WithSettlement(t *testing.T) {
	expenseRepo := &mockExpenseRepo{
		splits: []models.ExpenseSplitWithExpense{
			{
				UserID:    1,
				Amount:    500,
				PaidByID:  2,
				IsSettled: true,
			},
		},
	}

	groupRepo := &mockGroupRepo{
		requesterIsMember: true,
	}

	service := NewExpenseService(expenseRepo, groupRepo, nil)

	balances, err := service.CalculateBalances(1, 10)
	if err != nil {
		t.Fatal(err)
	}

	if balances[1] != 0 {
		t.Fatalf("expected user 1 balance 0, got %v", balances[1])
	}

	if balances[2] != 0 {
		t.Fatalf("expected user 2 balance 0, got %v", balances[2])
	}
}

func TestMarkAsPaid_Success(t *testing.T) {
	expenseRepo := &mockExpenseRepo{
		settleErr: nil,
	}

	groupRepo := &mockGroupRepo{
		requesterIsMember: true,
	}

	settleRepo := &mockSettlementRepo{}

	service := NewExpenseService(expenseRepo, groupRepo, settleRepo)

	req := dto.MarkPaidRequest{
		GroupID:   10,
		ExpenseID: 5,
	}

	err := service.MarkAsPaid(1, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestListExpensesWithShare_Success(t *testing.T) {

	expenseRepo := &mockExpenseRepo{
		expenseWithShare: []dto.ExpenseResponse{
			{
				ID:       1,
				Title:    "Dinner",
				Amount:   1000,
				PaidByID: 2,
				MyShare:  300,
			},
			{
				ID:       2,
				Title:    "Taxi",
				Amount:   400,
				PaidByID: 1,
				MyShare:  0,
			},
		},
	}

	groupRepo := &mockGroupRepo{
		requesterIsMember: true,
	}

	service := NewExpenseService(
		expenseRepo,
		groupRepo,
		nil,
	)

	expenses, err := service.ListExpensesWithShare(1, 10)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(expenses) != 2 {
		t.Fatalf("expected 2 expenses, got %d", len(expenses))
	}

	if expenses[0].MyShare != 300 {
		t.Fatalf("expected myShare 300, got %v", expenses[0].MyShare)
	}
}

func TestGetSettlementSuggestions_Success(t *testing.T) {

	groupRepo := &mockGroupRepo{
		requesterIsMember: true,
		members: []models.User{
			{ID: 1, Username: "pranav"},
			{ID: 2, Username: "raju"},
			{ID: 3, Username: "aman"},
		},
	}

	expenseRepo := &mockExpenseRepo{
		splits: []models.ExpenseSplitWithExpense{
			{UserID: 1, PaidByID: 1, Amount: 100},
			{UserID: 2, PaidByID: 1, Amount: 100},
			{UserID: 3, PaidByID: 1, Amount: 100},
		},
	}

	service := NewExpenseService(expenseRepo, groupRepo, nil)

	suggestions, err := service.GetSettlementSuggestions(1, 1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(suggestions) != 2 {
		t.Fatalf("expected 2 suggestions, got %d", len(suggestions))
	}

	if suggestions[0].FromUser != "raju" ||
		suggestions[0].ToUser != "pranav" ||
		suggestions[0].Amount != 100 {
		t.Fatalf("unexpected suggestion %v", suggestions[0])
	}

	if suggestions[1].FromUser != "aman" ||
		suggestions[1].ToUser != "pranav" ||
		suggestions[1].Amount != 100 {
		t.Fatalf("unexpected suggestion %v", suggestions[1])
	}
}
