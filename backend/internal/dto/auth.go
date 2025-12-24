package dto

type RegisterRequest struct {
	Username string  `json:"name"`
	Email    string  `json:"email"`
	Phone    *string `json:"phone"`
	Password string  `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
