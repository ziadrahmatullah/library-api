package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindBooks(context.Context) ([]models.Book, error)
	FindBooksByTitle(context.Context, string) ([]models.Book, error)
	FindBooksById(context.Context, uint) (*models.Book, error)
	NewBook(context.Context, models.Book) (*models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (b *bookRepository) FindBooks(ctx context.Context) (books []models.Book, err error) {
	b.db.WithContext(ctx).Exec("select pg_sleep(10)")
	err = b.db.WithContext(ctx).Preload("Author").Table("books").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *bookRepository) FindBooksByTitle(ctx context.Context, title string) (books []models.Book, err error) {
	err = b.db.WithContext(ctx).Preload("Author").Table("books").Where("title = ?", title).Find(&books).Error
	if err != nil {
		return nil, apperror.ErrFindBooksByTitleQuery
	}
	return books, nil
}

func (b *bookRepository) FindBooksById(ctx context.Context, id uint) (book *models.Book, err error) {
	result := b.db.WithContext(ctx).Table("books").Where("id = ?", id).Find(&book)
	if result.Error != nil {
		return nil, apperror.ErrFindBooksByTitleQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrBookNotFound
	}
	return book, nil
}

func (b *bookRepository) NewBook(ctx context.Context, book models.Book) (newBook *models.Book, err error) {
	err = b.db.WithContext(ctx).Table("books").Create(&book).Error
	if err != nil {
		return nil, apperror.ErrNewBookQuery
	}
	return &book, nil
}
