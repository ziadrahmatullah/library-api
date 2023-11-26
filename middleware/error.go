package middleware

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"github.com/gin-gonic/gin"
)

var statusCode = map[int]int{
	1: http.StatusBadRequest,
	2: http.StatusNotFound,
	3: http.StatusConflict,
}

func GetStatusCode(a apperror.HandlerErrType) int {
	return statusCode[int(a)]
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			var e apperror.Type
			switch {
			case errors.As(err.Err, &e):
				c.AbortWithStatusJSON(GetStatusCode(e.Type), gin.H{
					"error": err.Error(),
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
		}
	}
}
