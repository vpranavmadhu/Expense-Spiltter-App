package services

import (
	"errors"
	"esapp/internal/dto"
	"esapp/internal/repository"
	"esapp/models"
	"math"
)

type ExpenseService interface {
	CreateExpense(payerID uint, req dto.CreateExpenseRequest) error
	ListExpenses(requesterID, groupID uint) ([]models.Expense, error)
	CalculateBalances(requesterID, groupID uint) (map[uint]float64, error)
	MarkAsPaid(requesterID uint, req dto.MarkPaidRequest) error
	ListExpensesWithShare(groupID, requesterID uint) ([]dto.ExpenseResponse, error)
	GetSettlementSuggestions(groupID, requesterID uint) ([]dto.SettlementSuggestionResponse, error)
}

type expenseService struct {
	expenseRepo    repository.ExpenseRepository
	groupRepo      repository.GroupRepository
	settlementRepo repository.SettlementRepository
}

func NewExpenseService(expenseRepo repository.ExpenseRepository, groupRepo repository.GroupRepository, settlementRepo repository.SettlementRepository) ExpenseService {
	return &expenseService{expenseRepo: expenseRepo, groupRepo: groupRepo, settlementRepo: settlementRepo}
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

	var splits []models.ExpenseSplit

	if len(req.Splits) > 0 {
		total := 0.0
		memberSet := make(map[uint]bool)
		for _, m := range members {
			memberSet[m.ID] = true
		}

		for _, s := range req.Splits {
			if !memberSet[s.UserID] {
				return errors.New("Split user not in group")
			}

			total += s.Amount

			splits = append(splits, models.ExpenseSplit{
				ExpenseID: expense.ID,
				UserID:    s.UserID,
				Amount:    s.Amount,
			})
		}

		if math.Abs(total-req.Amount) > 0.01 {
			return errors.New("splits amount do not match total")
		}

	} else {
		//equal split
		splitAmount := req.Amount / float64(len(members))

		splits = make([]models.ExpenseSplit, 0)
		for _, m := range members {
			splits = append(splits, models.ExpenseSplit{
				ExpenseID: expense.ID,
				UserID:    m.ID,
				Amount:    splitAmount,
			})
		}
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
		return nil, errors.New("not authorized")
	}

	splits, err := s.expenseRepo.GetSplitsByGroupID(groupID)
	if err != nil {
		return nil, err
	}

	balances := make(map[uint]float64)

	for _, sp := range splits {

		if sp.IsSettled {
			continue
		}

		balances[sp.UserID] -= sp.Amount

		balances[sp.PaidByID] += sp.Amount
	}

	return balances, nil
}

func (s *expenseService) MarkAsPaid(requesterID uint, req dto.MarkPaidRequest) error {
	isMember, err := s.groupRepo.IsMember(req.GroupID, requesterID)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("not authorized")
	}

	payment := models.SettlementPayment{
		GroupID:    req.GroupID,
		FromUserID: requesterID,
		ToUserID:   req.ToUserID,
		Amount:     req.Amount,
	}

	if err := s.settlementRepo.Create(&payment); err != nil {
		return err
	}

	return s.expenseRepo.SettleExpenseSplit(req.ExpenseID, requesterID)

}

func (s *expenseService) ListExpensesWithShare(groupID, requesterID uint) ([]dto.ExpenseResponse, error) {

	isMember, err := s.groupRepo.IsMember(groupID, requesterID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("not authorized")
	}

	return s.expenseRepo.GetExpensesWithMyShare(groupID, requesterID)
}

func (s *expenseService) GetSettlementSuggestions(groupID, requesterID uint) ([]dto.SettlementSuggestionResponse, error) {

	isMember, err := s.groupRepo.IsMember(groupID, requesterID)
	if err != nil || !isMember {
		return nil, errors.New("not authorized")
	}

	balances, err := s.CalculateBalances(requesterID, groupID)
	if err != nil {
		return nil, err
	}

	members, err := s.groupRepo.GetMembersByGroupID(groupID)
	if err != nil {
		return nil, err
	}

	userMap := make(map[uint]string)
	for _, m := range members {
		userMap[m.ID] = m.Username
	}

	type pair struct {
		userID uint
		amount float64
	}

	var debtors []pair
	var creditors []pair

	for uid, bal := range balances {
		if bal < 0 {
			debtors = append(debtors, pair{uid, -bal})
		} else if bal > 0 {
			creditors = append(creditors, pair{uid, bal})
		}
	}

	var result []dto.SettlementSuggestionResponse

	i, j := 0, 0
	for i < len(debtors) && j < len(creditors) {
		d := &debtors[i]
		c := &creditors[j]

		amt := math.Min(d.amount, c.amount)

		result = append(result, dto.SettlementSuggestionResponse{
			FromUserID: d.userID,
			FromUser:   userMap[d.userID],
			ToUserID:   c.userID,
			ToUser:     userMap[c.userID],
			Amount:     math.Round(amt*100) / 100,
		})

		d.amount -= amt
		c.amount -= amt

		if d.amount == 0 {
			i++
		}
		if c.amount == 0 {
			j++
		}
	}

	return result, nil
}
