package handler

import (
	"net/http"
	"strconv"

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

func (h *BorrowingRecordHandler) BorrowBook(c *gin.Context) {
	var request dto.BorrowingRecordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		return
	}
	record := request.ToBorrowingRecord()
	createdBorrowingRecord, err := h.borrowingUsecase.BorrowBook(c.Request.Context(), record)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": dto.NewFromBorrowingRecord(createdBorrowingRecord),
	})
}

func (h *BorrowingRecordHandler) ReturnBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errPath := apperror.NewInvalidPathQueryParamError(err)
		_ = c.Error(errPath)
		return
	}
	br, err := h.borrowingUsecase.ReturnBook(c.Request.Context(), uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": dto.NewFromBorrowingRecord(br),
	})
}
