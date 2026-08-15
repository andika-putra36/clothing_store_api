package auth

import (
	"clothing_store_api/pkg/bcrypt"
	"clothing_store_api/pkg/jwt"
	"errors"
	"time"
)

type Service interface {
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) LogIn(input LoginRequest) (LoginResponse, error) {
	var loginCredential GetLoginCredentialResponse
	// fmt.Println(loginCredential)

	// Fetch data
	loginCredential, err := s.repository.GetLoginCredentials(input)
	if err != nil {
		return LoginResponse{}, err
	}

	// Compare password
	err = bcrypt.ComparePassword(loginCredential.PasswordHash, input.Password)
	if err != nil {
		return LoginResponse{}, err
	}

	// Generate access token
	accessToken, err := jwt.GenerateAccessToken(loginCredential.UserID, loginCredential.Email, loginCredential.RoleID)
	if err != nil {
		return LoginResponse{}, err
	}

	// Generate refresh token
	refreshToken, expiredAt, err := jwt.GenerateRefreshToken()
	if err != nil {
		return LoginResponse{}, err
	}

	// Delete refresh token in DB if exists
	err = s.repository.DeleteRefreshToken(loginCredential.UserID)
	if err != nil {
		return LoginResponse{}, err
	}

	// Store refresh token in DB
	err = s.repository.SaveRefreshToken(SaveRefreshTokenRequest{
		UserID:    loginCredential.UserID,
		Token:     refreshToken,
		ExpiredAt: expiredAt,
	})

	return LoginResponse{
		Token: Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s *service) RefreshToken(input RefreshTokenRequest) (RefreshTokenResponse, error) {
	tokenData, err := s.repository.GetRefreshToken(input.RefreshToken)
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	if time.Now().UTC().After(tokenData.ExpiredAt) {
		return RefreshTokenResponse{}, errors.New("Refresh token expired")
	}

	accessToken, err := jwt.GenerateAccessToken(tokenData.UserID, tokenData.Email, tokenData.RoleID)
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	newRefreshToken, expiredAt, err := jwt.GenerateRefreshToken()
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	err = s.repository.DeleteRefreshToken(tokenData.UserID)
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	err = s.repository.SaveRefreshToken(SaveRefreshTokenRequest{
		UserID:    tokenData.UserID,
		Token:     newRefreshToken,
		ExpiredAt: expiredAt,
	})
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	return RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
