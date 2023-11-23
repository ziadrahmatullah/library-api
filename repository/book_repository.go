package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type BookRepository interface{
	FindAllBooks() ([]models.Book, error)
}

type bookRepository struct{
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository{
	return &bookRepository{
		db: db,
	}
}

func (b *bookRepository) FindAllBooks() (books []models.Book, err error){
	err = b.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}