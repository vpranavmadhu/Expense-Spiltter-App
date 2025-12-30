package repository

import (
	"esapp/internal/dto"
	"esapp/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	CreateSplits(splits []models.ExpenseSplit) error
	GetExpensesByGroupID(groupID uint) ([]models.Expense, error)
	GetSplitsByGroupID(groupID uint) ([]models.ExpenseSplitWithExpense, error)
	GetExpensesWithMyShare(groupID uint, userID uint) ([]dto.ExpenseResponse, error)
	SettleExpenseSplit(expenseID, userID uint) error
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
	var splits []models.ExpenseSplitWithExpense

	err := r.db.Table("expense_splits").
		Select(`
			expense_splits.user_id,
			expense_splits.amount,
			expenses.paid_by_id,
			expense_splits.is_settled
		`).
		Joins("JOIN expenses ON expenses.id = expense_splits.expense_id").
		Where("expenses.group_id = ?", groupID).
		Scan(&splits).Error

	return splits, err
}

func (r *expenseRepository) GetExpensesWithMyShare(groupID uint, userID uint) ([]dto.ExpenseResponse, error) {

	var results []dto.ExpenseResponse

	err := r.db.Raw(`
		SELECT
			e.id,
			e.title,
			e.amount,
			e.paid_by_id AS paid_by_id,
			u.username AS paid_by_name,

			s.amount AS my_share,
			s.is_settled AS is_settled,

			COUNT(es.id) AS split_count,

			CASE
				WHEN e.paid_by_id = ?
				THEN SUM(es.amount) - COALESCE(ps.paid_sum, 0)
				ELSE 0
			END AS total_owed_to_me

		FROM expenses e

		JOIN users u ON u.id = e.paid_by_id
		JOIN expense_splits s ON s.expense_id = e.id AND s.user_id = ?
		JOIN expense_splits es ON es.expense_id = e.id

		LEFT JOIN (
			SELECT
				expense_id,
				SUM(amount) AS paid_sum
			FROM expense_splits
			WHERE is_settled = true
			GROUP BY expense_id
		) ps ON ps.expense_id = e.id

		WHERE e.group_id = ?

		GROUP BY
			e.id, u.username, s.amount, s.is_settled, ps.paid_sum
	`,
		userID, // for totalOwedToMe CASE
		userID, // my share
		groupID,
	).Scan(&results).Error

	return results, err
}

func (r *expenseRepository) SettleExpenseSplit(expenseID, userID uint) error {
	return r.db.Model(&models.ExpenseSplit{}).
		Where("expense_id = ? AND user_id = ?", expenseID, userID).
		Update("is_settled", true).Error
}
