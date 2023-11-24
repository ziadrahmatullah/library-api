package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"

type BorrowReq struct {
	UserId uint `json:"user_id" binding:"required"`
	BookId uint `json:"book_id" binding:"required"`
}

func (b *BorrowReq) ToModel(status string) models.BorrowingBook {
	return models.BorrowingBook{UserId: b.UserId, BookId: b.BookId, Status: status}
}
