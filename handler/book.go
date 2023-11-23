package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bookUsecase: bookUsecase,
	}
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	title := c.Query("title")
	var books []*entity.Book
	if title == "" {
		books = h.bookUsecase.GetAllBooks()
	} else {
		books = h.bookUsecase.FindBooksByTitle(title)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
