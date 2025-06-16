package internal

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateTkn(uid int, key string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": uid,
		"exp":     time.Now().Add(time.Hour * 24 * 365).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(key))
}
