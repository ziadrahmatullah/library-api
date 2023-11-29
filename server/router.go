package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/logger"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/middleware"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	BookHandler   *handler.BookHandler
	UserHandler   *handler.UserHandler
	BorrowHandler *handler.BorrowHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.ContextWithFallback = true

	router.Use(middleware.WithTimeout)
	router.Use(middleware.AuthorizeHandler())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.Logger(logger.NewLogger()))

	books := router.Group("/books")
	books.GET("", opts.BookHandler.HandleGetBooks)
	books.POST("", opts.BookHandler.HandleCreateBook)

	users := router.Group("/users")
	users.GET("", opts.UserHandler.HandleGetUsers)
	users.POST("/register", opts.UserHandler.HandleUserRegister)
	users.POST("/login", opts.UserHandler.HandleUserLogin)

	borrow := router.Group("/borrows")
	borrow.GET("", opts.BorrowHandler.HandleGetRecords)
	borrow.POST("", opts.BorrowHandler.HandleBorrowBook)
	borrow.PUT("", opts.BorrowHandler.HandleReturnBook)
	return router
}
