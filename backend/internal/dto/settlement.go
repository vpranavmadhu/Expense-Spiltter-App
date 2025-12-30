package dto

import "time"

type MarkPaidRequest struct {
	GroupID   uint    `json:"group_id" binding:"required"`
	ExpenseID uint    `json:"expense_id" binding:"required"`
	ToUserID  uint    `json:"to_user_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

type PaymentHistoryRespone struct {
	GroupName string    `json:"groupName"`
	FromUser  string    `json:"fromUser"`
	FromEmail string    `json:"fromEmail"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	Direction string    `json:"direction"` // paid or recieved
}
