package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"

type BorrowRes struct {
	Id     uint   `json:"id"`
	UserId uint   `json:"user_id"`
	BookId uint   `json:"book_id"`
	Status string `json:"status"`
}

func ToResponse(model *models.BorrowingBook) BorrowRes {
	return BorrowRes{Id: model.ID, UserId: model.UserId, BookId: model.BookId, Status: model.Status}
}
