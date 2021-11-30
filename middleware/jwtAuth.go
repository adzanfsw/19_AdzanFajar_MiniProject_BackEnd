package middleware

import (
	"justrun/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(ID int) (string, error) {

	payload := jwt.MapClaims{
		"user_id": ID,
		"role":    "admin",
		"exp":     time.Now().Add(time.Hour * 96).Unix(),
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	token, error := tokenWithClaims.SignedString([]byte(config.JwtSecret))

	return token, error
}
