package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/gin-gonic/gin"
)


type BorrowHandler struct{
	borrowUsecase usecase.BorrowUsecase
}

func NewBorrowHandler(bu usecase.BorrowUsecase) *BorrowHandler{
	return &BorrowHandler{
		borrowUsecase: bu,
	}
}

func (h *BorrowHandler) HandleBorrowBook(ctx *gin.Context){
	resp := dto.Response{}
	newBorrow := models.BorrowingBooks{}
	err := ctx.ShouldBindJSON(&newBorrow)
	if err != nil {
		resp.Message = apperror.ErrCannotBindJSON.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	borrow, err := h.borrowUsecase.BorrowBook(&newBorrow)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = borrow
	ctx.JSON(http.StatusOK, resp)
}