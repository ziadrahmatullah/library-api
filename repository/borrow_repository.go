package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type BorrowRepository interface{
	BorrowBook(*models.BorrowingBooks) (*models.BorrowingBooks, error)
}

type borrowRepository struct{
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository{
	return &borrowRepository{
		db: db,
	}
}

func (b *borrowRepository) BorrowBook(borrow *models.BorrowingBooks)(*models.BorrowingBooks, error){
	tx := b.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var book *models.Book
	err := tx.Table("books").Where("id = ?", borrow.BookId).Find(&book).Error
	if err != nil {
		return nil, err
	}
	if book == nil{
		return nil, apperror.ErrBookNotFound
	}
	if book.Quantity == 0 {
		return nil, apperror.ErrBookOutOfStock
	}
	err = tx.Table("books").Where("id = ?", book.ID).Update("quantity", gorm.Expr("quantity - ?", 1)).Error
	if err != nil {
		return nil, err
	}
	err = tx.Create(&borrow).Error
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return borrow, nil

}