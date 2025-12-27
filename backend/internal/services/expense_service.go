package services

import (
	"errors"
	"esapp/internal/dto"
	"esapp/internal/repository"
	"esapp/models"
)

type ExpenseService interface {
	CreateExpense(payerID uint, req dto.CreateExpenseRequest) error
	ListExpenses(requesterID, groupID uint) ([]models.Expense, error)
	CalculateBalances(requesterID, groupID uint) (map[uint]float64, error)
}

type expenseService struct {
	expenseRepo repository.ExpenseRepository
	groupRepo   repository.GroupRepository
}

func NewExpenseService(expenseRepo repository.ExpenseRepository, groupRepo repository.GroupRepository) ExpenseService {
	return &expenseService{expenseRepo: expenseRepo, groupRepo: groupRepo}
}

func (s *expenseService) CreateExpense(payerID uint, req dto.CreateExpenseRequest) error {

	isMember, err := s.groupRepo.IsMember(req.GroupID, payerID)
	if err != nil {
		return err
	}
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

func (s *expenseService) ListExpenses(requesterID, groupID uint) ([]models.Expense, error) {

	isMember, err := s.groupRepo.IsMember(groupID, requesterID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("not authorized")
	}

	return s.expenseRepo.GetExpensesByGroupID(groupID)
}

func (s *expenseService) CalculateBalances(requesterID, groupID uint) (map[uint]float64, error) {

	isMember, err := s.groupRepo.IsMember(groupID, requesterID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, err
	}

	splits, err := s.expenseRepo.GetSplitsByGroupID(groupID)
	if err != nil {
		return nil, err
	}

	balances := make(map[uint]float64)

	for _, sp := range splits {
		balances[sp.UserID] -= sp.Amount //user owes
		balances[sp.PaidByID] += sp.Amount
	}

	return balances, nil
}
