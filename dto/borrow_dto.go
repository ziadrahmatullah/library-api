package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"

type BorrowReq struct {
	UserId uint `json:"user_id" binding:"required"`
	BookId uint `json:"book_id" binding:"required"`
}

func (b *BorrowReq) ToBorrowModel() models.BorrowBook {
	return models.BorrowBook{
		UserId: b.UserId, 
		BookId: b.BookId, 
		Status: "not returned"}
}

type BorrowRes struct {
	ID     uint   `json:"id"`
	UserId uint   `json:"user_id"`
	BookId uint   `json:"book_id"`
	Status string `json:"status"`
}

func ToBorrowResponse(model *models.BorrowBook) BorrowRes {
	return BorrowRes{
		ID: model.ID, 
		UserId: model.UserId, 
		BookId: model.BookId, 
		Status: model.Status}
}
