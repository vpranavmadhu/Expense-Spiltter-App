package dto

type MarkPaidRequest struct {
	GroupID   uint `json:"group_id" binding:"required"`
	ExpenseID uint `json:"expense_id" binding:"required"`
}
