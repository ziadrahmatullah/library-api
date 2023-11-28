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
	Auth            *handler.AuthHandler
}

func New(handler Handlers) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	router.POST("/register", handler.Auth.Register)
	router.POST("/login", handler.Auth.Login)

	router.GET("/users", handler.User.GetAllUsers)
	router.GET("/books", handler.Book.GetAllBooks)

	if gin.Mode() != gin.DebugMode {
		router.Use(middleware.AuthHandler)
	}

	router.POST("/books", handler.Book.AddBook)

	router.POST("/borrowing-records", handler.BorrowingRecord.AddBorrowing)
	router.PUT("/borrowing-records/:id", handler.BorrowingRecord.ReturnBook)

	return router
}
