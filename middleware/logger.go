package middleware

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/logger"
	"github.com/gin-gonic/gin"
)

func Logger(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path

		ctx.Next()

		param := map[string]interface{}{
			"status_code": ctx.Writer.Status(),
			"method":      ctx.Request.Method,
			"latency":     time.Since(start),
			"path":        path,
		}

		if len(ctx.Errors) == 0 {
			log.Info(param)
		} else {
			errList := []error{}
			for _, err := range ctx.Errors {
				errList = append(errList, err)
			}

			if len(errList) > 0 {
				param["errors"] = errList
				log.Errorf("Invalid password", param)
			}
		}
	}

}
