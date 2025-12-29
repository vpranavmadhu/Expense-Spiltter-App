package models

type ExpenseSplitWithExpense struct {
	UserID    uint
	Amount    float64
	PaidByID  uint
	IsSettled bool
}
