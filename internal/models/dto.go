package models

type RegisterDTO struct {
	FirstName string  `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name"`
	Email     string  `json:"email" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
