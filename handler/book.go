package handler

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/dto"
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
	cl, err := getClause(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	title := c.Query("title")
	var books []*entity.Book
	if title == "" {
		books = h.bookUsecase.GetAllBooks(*cl)
	} else {
		books = h.bookUsecase.FindBooksByTitle(title)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var request dto.BookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": strings.Split(err.Error(), "\n"),
		})
		return
	}
	book := request.ToBook()

	createdBook, err := h.bookUsecase.AddBook(book)
	if err != nil {
		var e apperror.ErrAlreadyExist
		if errors.As(err, &e) {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": createdBook,
	})
}
