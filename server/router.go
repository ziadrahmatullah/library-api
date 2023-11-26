package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	BookHandler *handler.BookHandler
	UserHandler    *handler.UserHandler
	BorrowHandler  *handler.BorrowHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	books := router.Group("/books")
	books.GET("", opts.BookHandler.HandleGetBooks)
	books.POST("", opts.BookHandler.HandleCreateBook)

	users := router.Group("/users")
	users.GET("", opts.UserHandler.HandleGetUsers)

	borrow := router.Group("/borrows")
	borrow.GET("", opts.BorrowHandler.HandleGetRecords)
	borrow.POST("", opts.BorrowHandler.HandleBorrowBook)
	borrow.PUT("", opts.BorrowHandler.HandleReturnBook)
	return router
}
