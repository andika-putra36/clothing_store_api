package jwt

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID     int
	Email      string
	RoleID     int
	CustomerID int
	AdminID    int
	jwt.RegisteredClaims
}

func GetClaims(c *gin.Context) (*Claims, bool) {
	value, exists := c.Get("claims")
	if !exists {
		return nil, false
	}
	claims, ok := value.(*Claims)
	return claims, ok
}

func GenerateAccessToken(userID int, email string, roleID int, customerID int, adminID int) (string, error) {
	claims := Claims{
		UserID:     userID,
		Email:      email,
		RoleID:     roleID,
		CustomerID: customerID,
		AdminID:    adminID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func GenerateRefreshToken() (string, time.Time, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", time.Time{}, err
	}

	token := hex.EncodeToString(bytes)
	expiredAt := time.Now().UTC().Add(7 * 24 * time.Hour) // Expires in 7 Days

	return token, expiredAt, nil
}

func ValidateAccessToken(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}
