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

type LoginResponse struct {
	Token Token `json:"token"`
	User  User  `json:"user"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	UserID int    `json:"user_id"`
	RoleID int    `json:"role_id"`
	Email  string `json:"email"`
}
