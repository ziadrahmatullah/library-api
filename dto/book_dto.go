package dto

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
)

type BookReq struct {
	Title       string `json:"title" binding:"required,max=35" validate:"required,max=35"`
	Description string `json:"description" binding:"required" validate:"required"`
	Quantity    *int   `json:"quantity" binding:"required,min=0" validate:"required,min=0"`
	Cover       string `json:"cover,omitempty"`
	AuthorId    uint   `json:"author_id" binding:"required" validate:"required"`
}

func (b *BookReq) ToBookModel() models.Book {
	return models.Book{Title: b.Title, Description: b.Description, Quantity: uint(*b.Quantity), Cover: b.Cover, AuthorId: b.AuthorId}
}
