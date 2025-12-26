package repository

import (
	"esapp/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	CreateSplits(splits []models.ExpenseSplit) error
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
