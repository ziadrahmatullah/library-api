package migration

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	authors := []*entity.Author{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
	}
	books := []*entity.Book{
		{Title: "How to eat", Description: "Explain how to eat", Quantity: 2, Cover: "kertas", AuthorId: 1},
		{Title: "How to drink", Description: "Explain how to drink", Quantity: 3, Cover: "Kertas", AuthorId: 2},
	}
	db.Create(authors)
	db.Create(books)
}
