package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindBooks() ([]models.Book, error)
	FindBooksByTitle(string) ([]models.Book, error)
	NewBook(*models.Book) (*models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

type BookParameter struct {
	Title string
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (b *bookRepository) FindBooks() (books []models.Book, err error) {
	table := b.db.Preload("Author").Table("books")
	err = table.Find(&books).Error
	if err != nil {
		return nil, apperror.ErrFindBooksQuery
	}
	return books, nil
}

func (b *bookRepository) FindBooksByTitle(title string) (books []models.Book, err error) {
	table := b.db.Preload("Author").Table("books")
	err = table.Where("title = ?", title).Find(&books).Error
	if err != nil {
		return nil, apperror.ErrFindBooksByTitleQuery
	}
	return books, nil
}

func (b *bookRepository) NewBook(book *models.Book) (newBook *models.Book, err error) {
	err = b.db.Table("books").Create(&book).Error
	if err != nil {
		return nil, apperror.ErrNewBookQuery
	}
	return book, nil
}
