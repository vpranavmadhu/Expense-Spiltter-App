package services

import (
	"errors"
	"esapp/internal/dto"
	"esapp/internal/repository"
	"esapp/models"
)

type ExpenseService interface {
	CreateExpense(payerID uint, req dto.CreateExpenseRequest) error
}

type expenseService struct {
	expenseRepo repository.ExpenseRepository
	groupRepo   repository.GroupRepository
}

func NewExpenseService(expenseRepo repository.ExpenseRepository, groupRepo repository.GroupRepository) ExpenseService {
	return &expenseService{expenseRepo: expenseRepo, groupRepo: groupRepo}
}

func (s *expenseService) CreateExpense(payerID uint, req dto.CreateExpenseRequest) error {

	isMember, _ := s.groupRepo.IsMember(req.GroupID, payerID)
	if !isMember {
		return errors.New("not authorized")
	}

	members, err := s.groupRepo.GetMembersByGroupID(req.GroupID)
	if err != nil || len(members) == 0 {
		return errors.New("no group members")
	}

	expense := models.Expense{
		GroupID:  req.GroupID,
		PaidByID: payerID,
		Title:    req.Title,
		Amount:   req.Amount,
	}

	if err := s.expenseRepo.CreateExpense(&expense); err != nil {
		return err
	}

	//equal split
	splitAmount := req.Amount / float64(len(members))

	splits := make([]models.ExpenseSplit, 0)
	for _, m := range members {
		splits = append(splits, models.ExpenseSplit{
			ExpenseID: expense.ID,
			UserID:    m.ID,
			Amount:    splitAmount,
		})
	}
	return s.expenseRepo.CreateSplits(splits)

}
