package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
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
	q, err := getQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	title := c.Query("title")
	quantity := c.Query("qty")
	description := c.Query("desc")

	conditions := []*valueobject.Condition{
		valueobject.NewCondition("title", valueobject.Ilike, title),
		valueobject.NewCondition("description", valueobject.Ilike, description),
		valueobject.NewCondition("quantity", valueobject.Equal, quantity),
	}
	q.Conditions = filterCondition(conditions)
	q.With = []string{"Author"}
	var books []*entity.Book
	books, err = h.bookUsecase.GetAllBooks(c.Request.Context(), q)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": dto.NewFromBooks(books),
	})
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var request dto.BookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		return
	}
	book := request.ToBook()

	createdBook, err := h.bookUsecase.AddBook(c.Request.Context(), book)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": dto.NewFromBook(createdBook),
	})
}
