package repository

import (
	"esapp/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	CreateSplits(splits []models.ExpenseSplit) error
	GetExpensesByGroupID(groupID uint) ([]models.Expense, error)
	GetSplitsByGroupID(groupID uint) ([]models.ExpenseSplitWithExpense, error)
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) CreateExpense(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) CreateSplits(splits []models.ExpenseSplit) error {
	return r.db.Create(&splits).Error
}

func (r *expenseRepository) GetExpensesByGroupID(groupID uint) ([]models.Expense, error) {
	var expenses []models.Expense
	err := r.db.Where("group_id =?", groupID).Order("created_at DESC").Find(&expenses).Error
	return expenses, err
}

func (r *expenseRepository) GetSplitsByGroupID(groupID uint) ([]models.ExpenseSplitWithExpense, error) {
	var result []models.ExpenseSplitWithExpense

	err := r.db.Raw(`
		SELECT 
			es.user_id,
			es.amount,
			e.paid_by_id
		FROM expense_splits es
		JOIN expenses e ON e.id = es.expense_id
		WHERE e.group_id = ?
	`, groupID).Scan(&result).Error

	return result, err

}
