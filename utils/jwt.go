package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, nationalID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "change-this-secret"
	}

	claims := jwt.MapClaims{
		"user_id":     userID,
		"national_id": nationalID,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
		"iat":         time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}