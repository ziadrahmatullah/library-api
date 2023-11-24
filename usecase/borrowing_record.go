package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
)

type BorrowingRecordUsecase interface {
	AddBorrowingRecord(br *entity.BorrowingRecords) (*entity.BorrowingRecords, error)
}

type borrowingRecordUsecase struct {
	borrowingRepo repository.BorrowingRecordRepository
}

func NewBorrowingRecordUsecase(borrowingRepo repository.BorrowingRecordRepository) BorrowingRecordUsecase {
	return &borrowingRecordUsecase{
		borrowingRepo: borrowingRepo,
	}
}

func (u *borrowingRecordUsecase) AddBorrowingRecord(br *entity.BorrowingRecords) (*entity.BorrowingRecords, error) {
	return nil, nil
}
