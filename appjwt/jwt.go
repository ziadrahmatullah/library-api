package appjwt

import (
	"errors"
	"os"
	"strconv"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Id    uint
	Email string
}

type Jwt interface {
	GenerateToken(user *entity.User) (string, error)
	ValidateToken(tokenString string) (*CustomClaims, error)
}

type jwtImpl struct {
	secretKey []byte
}

func NewJwt() Jwt {
	key := os.Getenv("SECRET")
	return &jwtImpl{
		secretKey: []byte(key),
	}
}

func (j *jwtImpl) GenerateToken(user *entity.User) (string, error) {
	userId := strconv.Itoa(int(user.Id))
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "LIBRARY",
			Subject:   userId,
		},
		Id: user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(j.secretKey)
	return signedString, err
}

func (j *jwtImpl) ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid claims type")
}
