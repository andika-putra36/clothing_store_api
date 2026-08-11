package auth

import "time"

type GetLoginCredentialResponse struct {
	UserID       int    `json:"user_id"`
	RoleID       int    `json:"role_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type GetRefreshTokenResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
