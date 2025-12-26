package dto

type CreateExpenseRequest struct {
	GroupID uint    `json:"group_id" binding:"required"`
	Title   string  `json:"title" binding:"required"`
	Amount  float64 `json:"amount" binding:"required,gt=0"`
}
