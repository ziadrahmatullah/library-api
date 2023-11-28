package middleware

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func WithTimeout() gin.HandlerFunc {
	var defaultTimeout = 5
	timeoutString := os.Getenv("TIMEOUT")
	timeout, err := strconv.Atoi(timeoutString)
	if err != nil {
		timeout = defaultTimeout
	}
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(timeout)*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
