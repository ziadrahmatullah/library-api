package middleware

import (
	"errors"
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0].Err

		var clientError *apperror.ClientError
		var validationError validator.ValidationErrors

		isClientError := false

		if errors.As(err, &clientError) {
			isClientError = true
			err = clientError.UnWrap()
		}

		switch {
		case errors.As(err, &validationError):
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": strings.Split(validationError.Error(), "\n"),
			})
		case isClientError:
			c.AbortWithStatusJSON(clientError.GetCode(), gin.H{
				"error": clientError.Error(),
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
}
