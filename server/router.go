package server

import (
	"net/http"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
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

	router.Use(middleware.ErrorHandler())
	router.Use(middleware.WithTimeout)

	router.GET("/hello", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"data": "hello world",
		})
	})

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
