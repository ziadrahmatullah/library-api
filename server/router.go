package server

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	ProductHandler *handler.BookHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	product := router.Group("/books")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	product.GET("/", opts.ProductHandler.HandleGetAllBooks)
	return router
}
