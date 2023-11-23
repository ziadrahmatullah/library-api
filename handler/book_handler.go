package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/gin-gonic/gin"
)

type BookHandler struct{
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(pu usecase.BookUsecase) *BookHandler{
	return &BookHandler{
		bookUsecase: pu,
	}
}

func (h *BookHandler) HandleGetAllBooks(ctx *gin.Context){
	resp := dto.Response{}
	products, err := h.bookUsecase.GetAllBooks()
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = products
	ctx.JSON(http.StatusOK, resp)
}