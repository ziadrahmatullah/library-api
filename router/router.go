package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Book            *handler.BookHandler
	User            *handler.UserHandler
	BorrowingRecord *handler.BorrowingRecordHandler
}

func New(handler Handlers) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandler())
	
	router.GET("/books", handler.Book.GetAllBooks)
	router.POST("/books", handler.Book.AddBook)

	router.GET("/users", handler.User.GetAllUsers)

	router.POST("/borrowing-records", handler.BorrowingRecord.AddBorrowing)

	return router
}
