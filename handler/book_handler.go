package handler

import (
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

func NewBookHandler(bu usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bookUsecase: bu,
	}
}

func (h *BookHandler) HandleGetBooks(ctx *gin.Context) {
	resp := dto.Response{}
	title := ctx.Query("title")
	var books []models.Book
	var err error
	if title != "" {
		books, err = h.bookUsecase.GetBooksByTitle(ctx, title)
	} else {
		books, err = h.bookUsecase.GetAllBooks(ctx)
	}
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = books
	ctx.JSON(http.StatusOK, resp)
}

func (h *BookHandler) HandleCreateBook(ctx *gin.Context) {
	resp := dto.Response{}
	newBook := dto.BookReq{}
	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	book, err := h.bookUsecase.CreateBook(ctx, newBook.ToBookModel())
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = book
	ctx.JSON(http.StatusOK, resp)
}
