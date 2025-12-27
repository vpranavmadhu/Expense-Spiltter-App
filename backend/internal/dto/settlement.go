package dto

type MarkPaidRequest struct {
	GroupID uint    `json:"group_id" binding:"required"`
	ToUser  uint    `json:"to_user_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required,gt=0"`
}
