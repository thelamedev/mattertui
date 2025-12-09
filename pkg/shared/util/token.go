package util

import (
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
