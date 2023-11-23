package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	ProductHandler *handler.BookHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	books := router.Group("/books")
	books.GET("", opts.ProductHandler.HandleGetBooks)
	books.POST("", opts.ProductHandler.HandleCreateBook)
	return router
}
