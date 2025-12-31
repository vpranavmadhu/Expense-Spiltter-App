package dto

import "time"

type CreateGroupRequest struct {
	Name string `json:"name" binding:"required"`
}

type AddMemberRequest struct {
	Email string `json:"email" binding:"required"`
}

type GroupResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	CreatorName string    `json:"creator_name"`
	CreatedAt   time.Time `json:"created_at"`
	Members     []Member  `json:"members"`
}

type Member struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
