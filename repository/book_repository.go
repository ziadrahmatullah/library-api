package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository interface {
	FindBooks() ([]models.Book, error)
	FindBooksByTitle(string) ([]models.Book, error)
	FindBooksById(uint) (*models.Book, error)
	NewBook(models.Book) (*models.Book, error)
	IncreaseBookQty(uint) error
	DecreaseBookQty(uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (b *bookRepository) FindBooks() (books []models.Book, err error) {
	err = b.db.Preload("Author").Table("books").Find(&books).Error
	if err != nil {
		return nil, apperror.ErrFindBooksQuery
	}
	return books, nil
}

func (b *bookRepository) FindBooksByTitle(title string) (books []models.Book, err error) {
	err = b.db.Preload("Author").Table("books").Where("title = ?", title).Find(&books).Error
	if err != nil {
		return nil, apperror.ErrFindBooksByTitleQuery
	}
	return books, nil
}

func (b *bookRepository) FindBooksById(id uint) (book *models.Book, err error) {
	result := b.db.Table("books").Where("id = ?", id).Find(&book)
	if result.Error != nil {
		return nil, apperror.ErrFindBooksByTitleQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrBookNotFound
	}
	return book, nil
}

func (b *bookRepository) NewBook(book models.Book) (newBook *models.Book, err error) {
	err = b.db.Table("books").Create(&book).Error
	if err != nil {
		return nil, apperror.ErrNewBookQuery
	}
	return &book, nil
}

func (b *bookRepository) DecreaseBookQty(id uint) (err error) {
	tx := b.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return
	}
	err = tx.Table("books").Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).Update("quantity", gorm.Expr("quantity - ?", 1)).Error
	if err != nil {
		return apperror.ErrUpdateBookQtyQuery
	}
	err = tx.Commit().Error
	if err != nil {
		return apperror.ErrTxCommit
	}
	return nil
}

func (b *bookRepository) IncreaseBookQty(id uint) (err error) {
	tx := b.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return
	}
	err = tx.Table("books").Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).Update("quantity", gorm.Expr("quantity + ?", 1)).Error
	if err != nil {
		return apperror.ErrUpdateBookQtyQuery
	}
	err = tx.Commit().Error
	if err != nil {
		return apperror.ErrTxCommit
	}
	return nil
}
