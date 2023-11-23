package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Book *handler.BookHandler
}

func New(handler Handlers) *gin.Engine {
	router := gin.Default()

	router.GET("/books", handler.Book.GetAllBooks)
	router.POST("/books", handler.Book.AddBook)

	return router
}
