package middleware

import (
	"context"
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/appjwt"
	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	token, err := extractBearerToken(bearerToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	jwt := appjwt.NewJwt()
	claims, err := jwt.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": apperror.ErrInvalidToken{}.Error(),
		})
		return
	}
	newContext := context.WithValue(c.Request.Context(), "userId", claims.Id)
	c.Request = c.Request.WithContext(newContext)
	c.Next()
}

func extractBearerToken(bearerToken string) (string, error) {
	if bearerToken == "" {
		return "", apperror.ErrMissingToken{}
	}
	token := strings.Split(bearerToken, " ")
	if len(token) != 2 {
		return "", apperror.ErrInvalidToken{}
	}
	return token[1], nil
}
