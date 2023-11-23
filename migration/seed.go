package migration

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	books := []*entity.Book{
		{Title: "How to eat", Description: "Explain how to eat", Quantity: 2, Cover: "kertas"},
	}
	db.Create(books)
}
