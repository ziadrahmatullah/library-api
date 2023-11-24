package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/gin-gonic/gin"
)

type BorrowHandler struct {
	borrowUsecase usecase.BorrowUsecase
}

func NewBorrowHandler(bu usecase.BorrowUsecase) *BorrowHandler {
	return &BorrowHandler{
		borrowUsecase: bu,
	}
}

func (h *BorrowHandler) HandleBorrowBook(ctx *gin.Context) {
	resp := dto.Response{}
	newBorrow := dto.BorrowReq{}
	err := ctx.ShouldBindJSON(&newBorrow)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	borrowModel := newBorrow.ToModel("not complete")
	borrow, err := h.borrowUsecase.BorrowBook(&borrowModel)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	borrowRespond := dto.ToResponse(borrow)
	resp.Data = borrowRespond
	ctx.JSON(http.StatusOK, resp)
}
