package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"

type BookRequest struct {
	Title       string `json:"title" binding:"required,max=35"`
	Description string `json:"description" binding:"required"`
	Quantity    *int   `json:"quantity" binding:"required,min=0"`
	Cover       entity.Cover
	AuthorId    uint `json:"author_id" binding:"required"`
}

func (r BookRequest) ToBook() *entity.Book {
	return &entity.Book{
		Title:       r.Title,
		Description: r.Description,
		Quantity:    *r.Quantity,
		Cover:       r.Cover,
		AuthorId:    r.AuthorId,
	}
}
