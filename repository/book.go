package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() []*entity.Book
	FindByTitle(name string) []*entity.Book
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

func (r *bookRepository) FindByTitle(title string) []*entity.Book {
	var books []*entity.Book
	r.db.Where("title ILIKE ?", "%"+title+"%").Find(&books)
	return books
}
