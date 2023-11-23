package handler

import (
	"log"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(pu usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bookUsecase: pu,
	}
}

func (h *BookHandler) HandleGetBooks(ctx *gin.Context) {
	resp := dto.Response{}
	title := ctx.Query("title")
	var books []models.Book
	var err error
	if title != "" {
		books, err = h.bookUsecase.GetBooksByTitle(title)
	} else {
		books, err = h.bookUsecase.GetAllBooks()
	}
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = books
	ctx.JSON(http.StatusOK, resp)
}

func (h *BookHandler) HandleCreateBook(ctx *gin.Context) {
	resp := dto.Response{}
	newBook := models.Book{}
	err := ctx.ShouldBindJSON(&newBook)
	log.Print(newBook)
	if err != nil {
		resp.Message = apperror.ErrCannotBindJSON.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	book, err := h.bookUsecase.CreateBook(&newBook)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = book
	ctx.JSON(http.StatusOK, resp)

}
