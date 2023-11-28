package middleware

import (
	"context"
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				c.AbortWithStatusJSON(http.StatusGatewayTimeout, dto.Response{Message: "request timeout"})
				return
			}
			switch err.Err.(type) {
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
	}
}
