package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BorrowRepository interface {
	NewBorrow(context.Context, models.BorrowBook) (*models.BorrowBook, error)
	FindBorrows(context.Context) ([]models.BorrowBook, error)
	FindBorrow(context.Context, models.BorrowBook) (uint, error)
	UpdateBorrowStatus(context.Context, uint) (*models.BorrowBook, error)
}

type borrowRepository struct {
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository {
	return &borrowRepository{
		db: db,
	}
}

func (b *borrowRepository) NewBorrow(ctx context.Context, borrow models.BorrowBook) (newBorrow *models.BorrowBook, err error) {
	tx := b.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.Table("books").
		Where("id = ?", borrow.BookId).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Update("quantity", gorm.Expr("quantity - ?", 1)).Error
	if err != nil {
		return nil, apperror.ErrUpdateQty
	}
	err = tx.Table("borrowing_books").Create(&borrow).Error
	if err != nil {
		return nil, apperror.ErrNewBorrowQuery
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, apperror.ErrTxCommit
	}
	return &borrow, nil
}

func (b *borrowRepository) FindBorrows(ctx context.Context) (borrows []models.BorrowBook, err error) {	
	err = b.db.WithContext(ctx).Preload("User").Preload("Book").Table("borrowing_books").Find(&borrows).Error
	if err != nil {
		return nil, apperror.ErrFindBorrowsQuery
	}
	return borrows, nil
}

func (b *borrowRepository) FindBorrow(ctx context.Context, borrow models.BorrowBook) (id uint,err error) {
	result := b.db.WithContext(ctx).Table("borrowing_books").
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

func (b *borrowRepository) UpdateBorrowStatus(ctx context.Context, id uint) (updatedBorrow *models.BorrowBook, err error) {
	tx := b.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var borrow models.BorrowBook
	err = tx.Table("borrowing_books").
		Where("id = ?", id).Update("status", "returned").
		Scan(&borrow).Error
	if err != nil {
		return nil, apperror.ErrUpdateStatus
	}
	err = tx.Table("books").
		Where("id = ?", borrow.BookId).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Update("quantity", gorm.Expr("quantity + ?", 1)).Error
	if err != nil {
		return nil, apperror.ErrUpdateQty
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, apperror.ErrTxCommit
	}
	return &borrow, nil
}
