package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type BorrowRepository interface {
	NewBorrow(models.BorrowBook) (*models.BorrowBook, error)
	FindBorrows() ([]models.BorrowBook, error)
	FindBorrow(models.BorrowBook) (uint, error)
	UpdateBorrowStatus(uint) (*models.BorrowBook, error)
}

type borrowRepository struct {
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository {
	return &borrowRepository{
		db: db,
	}
}

func (b *borrowRepository) NewBorrow(borrow models.BorrowBook) (newBorrow *models.BorrowBook, err error) {
	err = b.db.Table("borrowing_books").Create(&borrow).Error
	if err != nil {
		return nil, apperror.ErrNewBorrowQuery
	}
	return &borrow, nil
}

func (b *borrowRepository) FindBorrows() (borrows []models.BorrowBook, err error) {
	err = b.db.Preload("User").Preload("Book").Table("borrowing_books").Find(&borrows).Error
	if err != nil {
		return nil, apperror.ErrFindBooksQuery
	}
	return borrows, nil
}

func (b *borrowRepository) FindBorrow(borrow models.BorrowBook) (id uint,err error) {
	result := b.db.Table("borrowing_books").
		Where("user_id = ? AND book_id = ? AND status = ?", borrow.UserId, borrow.BookId, "not returned").
		Order("id desc").
		Pluck("id", &id)
	if result.Error != nil {
		return 0, apperror.ErrFindBorrowQuery
	}
	if result.RowsAffected == 0 {
		return 0, apperror.ErrBorrowRecordNotFound
	}
	return id, nil
}

func (b *borrowRepository) UpdateBorrowStatus(id uint) (updatedBorrow *models.BorrowBook, err error) {
	var borrow models.BorrowBook
	err = b.db.Table("borrowing_books").
		Where("id = ?", id).Update("status", "returned").
		Scan(&borrow).Error
	if err != nil {
		return nil, apperror.ErrNewBorrowQuery
	}
	return &borrow, nil
}