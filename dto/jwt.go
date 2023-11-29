package dto

import (
	"os"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	ID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(claims JwtClaims) (string, error) {
	claims.Issuer = "authToken"
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, apperror.ErrInvalidPassword
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
