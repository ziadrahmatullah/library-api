package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
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

func (h *BorrowHandler) HandleGetRecords(ctx *gin.Context) {
	resp := dto.Response{}
	records, err := h.borrowUsecase.GetAllRecords(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = records
	ctx.JSON(http.StatusOK, resp)
}

func (h *BorrowHandler) HandleBorrowBook(ctx *gin.Context) {
	resp := dto.Response{}
	newBorrow := dto.BorrowReq{}
	err := ctx.ShouldBindJSON(&newBorrow)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	reqContext := dto.CreateContext(ctx)
	borrowModel := newBorrow.ToBorrowModel(reqContext.UserID)
	borrow, err := h.borrowUsecase.BorrowBook(ctx, borrowModel)
	if err != nil {
		ctx.Error(err)
		return
	}
	borrowRespond := dto.ToBorrowResponse(borrow)
	resp.Data = borrowRespond
	ctx.JSON(http.StatusOK, resp)
}

func (h *BorrowHandler) HandleReturnBook(ctx *gin.Context) {
	resp := dto.Response{}
	borrowRecord := dto.BorrowReq{}
	err := ctx.ShouldBindJSON(&borrowRecord)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	reqContext := dto.CreateContext(ctx)
	borrowModel := borrowRecord.ToBorrowModel(reqContext.UserID)
	borrow, err := h.borrowUsecase.ReturnBook(ctx, borrowModel)
	if err != nil {
		ctx.Error(err)
		return
	}
	borrowRespond := dto.ToBorrowResponse(borrow)
	resp.Data = borrowRespond
	ctx.JSON(http.StatusOK, resp)
}
