package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BorrowRepository interface {
	BorrowBook(*models.BorrowingBook) (*models.BorrowingBook, error)
}

type borrowRepository struct {
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository {
	return &borrowRepository{
		db: db,
	}
}

func (b *borrowRepository) BorrowBook(borrow *models.BorrowingBook) (*models.BorrowingBook, error) {
	// tx := b.db.Begin()
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// }()
	// var book *models.Book
	// err := tx.Table("books").Where("id = ?", borrow.BookId).Find(&book).Error
	// if err != nil {
	// 	return nil, err
	// }
	// if book == nil {
	// 	return nil, apperror.ErrBookNotFound
	// }
	// if book.Quantity == 0 {
	// 	return nil, apperror.ErrBookOutOfStock
	// }
	// err = tx.Table("books").Where("id = ?", book.ID).Update("quantity", gorm.Expr("quantity - ?", 1)).Error
	// if err != nil {
	// 	return nil, err
	// }
	// err = tx.Create(&borrow).Error
	// if err != nil {
	// 	return nil, err
	// }
	// tx.Commit()
	// return borrow, nil
	bookId := borrow.BookId
	err := b.db.Transaction(func(tx *gorm.DB) error {
		book := models.Book{}
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&book, bookId).Error
		// err := tx.First(&book, bookId).Error
		if err != nil {
			return apperror.ErrBookNotFound
		}
		if book.Quantity == 0 {
			return apperror.ErrBookOutOfStock
		}
		err = tx.Create(&borrow).Error
		if err != nil {
			return err
		}
		book.Quantity = book.Quantity - uint(1)
		err = tx.Save(&book).Error
		if err != nil {
			return err
		}
		return nil
	})
	return borrow, err

}
