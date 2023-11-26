package handler

import (
	"log"
	"net/http"
	"strings"

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
	var request dto.BorrowingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": strings.Split(err.Error(), "\n"),
		})
		return
	}
	record := request.ToBorrowingRecord()
	createdBorrowingRecord, err := h.borrowingUsecase.AddBorrowingRecord(c, record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": createdBorrowingRecord,
	})
}
