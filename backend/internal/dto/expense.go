package dto

type SplitInput struct {
	UserID uint    `json:"user_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type CreateExpenseRequest struct {
	GroupID uint         `json:"group_id" binding:"required"`
	Title   string       `json:"title" binding:"required"`
	Amount  float64      `json:"amount" binding:"required,gt=0"`
	Splits  []SplitInput `json:"splits"`
}

type ExpenseResponse struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title"`
	Amount     float64 `json:"amount"`
	PaidByID   uint    `json:"paidById"`
	PaidByName string  `json:"paidByName"`
	MyShare    float64 `json:"myShare"`
	IsSettled  bool    `json:"isSettled"`
}
