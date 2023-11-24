package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"

type BorrowingRequest struct {
	UserId uint `json:"user_id" binding:"required,number"`
	BookId uint `json:"book_id" binding:"required,number"`
}

func (br BorrowingRequest) ToBorrowingRecord() *entity.BorrowingRecords {
	return &entity.BorrowingRecords{
		UserId: br.UserId,
		BookId: br.BookId,
	}
}
