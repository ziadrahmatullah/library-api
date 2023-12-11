package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func WithTimeout(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func GrpcTimeout(){
	// log.Info().Msg("stopping server")
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()
	// server.GracefulStop()
	// <-ctx.Done()
}