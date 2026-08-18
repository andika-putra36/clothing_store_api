package auth

import "time"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SaveRefreshTokenRequest struct {
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RegisterCustomerRequest struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type RegisterAdminRequest struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
