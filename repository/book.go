package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() []*entity.Book
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) FindAll() []*entity.Book {
	var books []*entity.Book
	r.db.Find(&books)
	return books
}
