package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
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

func (r *borrowingRecordRepository) Create(ctx context.Context, br *entity.BorrowingRecords) (*entity.BorrowingRecords, error) {
	br.Status = 1
	br, err := r.baseRepository.Create(ctx, br)
	if err != nil {
		return nil, err
	}
	return br, nil
}
