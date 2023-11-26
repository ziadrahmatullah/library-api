package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"github.com/gin-gonic/gin"
)

type BorrowingRecordHandler struct {
	borrowingUsecase usecase.BorrowingRecordUsecase
}

func NewBorrowingRecordHandler(borrowingRecord usecase.BorrowingRecordUsecase) *BorrowingRecordHandler {
	return &BorrowingRecordHandler{
		borrowingUsecase: borrowingRecord,
	}
}

func (h *BorrowingRecordHandler) AddBorrowing(c *gin.Context) {
	var request dto.BorrowingRecordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(apperror.ErrBinding{ErrBinding: err})
		return
	}
	record := request.ToBorrowingRecord()
	createdBorrowingRecord, err := h.borrowingUsecase.AddBorrowingRecord(c, record)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": dto.NewFromBorrowingRecord(createdBorrowingRecord),
	})
}
