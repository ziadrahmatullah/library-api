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
	users := []*entity.User{
		{Name: "Alice", Email: "alice@example.com", Phone: "085346727162"},
		{Name: "Bob", Email: "bob@example.com", Phone: "085212819384"},
		{Name: "Charlie", Email: "charlie@example.com", Phone: "081394839283"},
	}
	borrowingRecords := []*entity.BorrowingRecords{
		{UserId: 1, BookId: 1, Status: 1},
		{UserId: 1, BookId: 2, Status: 1},
		{UserId: 2, BookId: 1, Status: 1},
		{UserId: 3, BookId: 2, Status: 1},
	}
	db.Create(authors)
	db.Create(books)
	db.Create(users)
	db.Create(borrowingRecords)
}
