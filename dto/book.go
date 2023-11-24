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

type BookResponse struct {
	Id          uint            `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Quantity    int             `json:"quantity"`
	Cover       entity.Cover    `json:"cover"`
	Author      *AuthorResponse `json:"author,omitempty"`
}

func NewFromBook(book *entity.Book) *BookResponse {
	var authorResponse *AuthorResponse
	if book.Author != nil {
		authorResponse = NewFromAuthor(book.Author)
	}
	return &BookResponse{
		Id:          book.Id,
		Title:       book.Title,
		Description: book.Description,
		Quantity:    book.Quantity,
		Cover:       book.Cover,
		Author:      authorResponse,
	}
}

func NewFromBooks(books []*entity.Book) []*BookResponse {
	return newResponsesFromEntities(books, NewFromBook)
}
