package dto

type CreateGroupRequest struct {
	Name string `json:"name" binding:"required"`
}

type AddMemberRequest struct {
	Email string `json:"email" binding:"required"`
}
