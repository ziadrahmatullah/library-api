package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"

type BorrowingRecordRequest struct {
	UserId uint `json:"user_id" binding:"required,number"`
	BookId uint `json:"book_id" binding:"required,number"`
}

func (br BorrowingRecordRequest) ToBorrowingRecord() *entity.BorrowingRecords {
	return &entity.BorrowingRecords{
		UserId: br.UserId,
		BookId: br.BookId,
	}
}

type BorrowingRecordResponse struct {
	Id     uint `json:"id"`
	UserId uint `json:"user_id"`
	BookId uint `json:"book_id"`
}

func NewFromBorrowingRecord(br *entity.BorrowingRecords) *BorrowingRecordResponse {
	return &BorrowingRecordResponse{
		Id:     br.Id,
		UserId: br.UserId,
		BookId: br.BookId,
	}
}
