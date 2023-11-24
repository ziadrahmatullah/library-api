package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BorrowingRecordRepository interface {
	BaseRepository[entity.BorrowingRecords]
}

type borrowingRecordRepository struct {
	*baseRepository[entity.BorrowingRecords]
	db *gorm.DB
}

func NewBorrowingRecordsRepository(db *gorm.DB) BorrowingRecordRepository {
	return &borrowingRecordRepository{
		db:             db,
		baseRepository: &baseRepository[entity.BorrowingRecords]{db: db},
	}
}

func (r *borrowingRecordRepository) Create(br *entity.BorrowingRecords) (*entity.BorrowingRecords, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		book := &entity.Book{}
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id", br.BookId).
			First(&book).
			Error
		if err != nil {
			return apperror.ErrAlreadyExist{
				Resource: "book",
			}
		}
		if book.Quantity < 1 {
			return apperror.ErrNotFound{}
		}
		err = tx.Model(book).Update("quantity", gorm.Expr("quantity-1")).Error
		if err != nil {
			return err
		}
		br.Status = 1
		err = tx.Create(br).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return br, nil
}
