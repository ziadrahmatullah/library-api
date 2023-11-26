package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BorrowingRecordUsecase interface {
	AddBorrowingRecord(ctx context.Context, br *entity.BorrowingRecords) (*entity.BorrowingRecords, error)
}

type borrowingRecordUsecase struct {
	borrowingRepo repository.BorrowingRecordRepository
	bookRepo      repository.BookRepository
}

func NewBorrowingRecordUsecase(borrowingRepo repository.BorrowingRecordRepository, bookRepo repository.BookRepository) BorrowingRecordUsecase {
	return &borrowingRecordUsecase{
		borrowingRepo: borrowingRepo,
		bookRepo:      bookRepo,
	}
}

func (u *borrowingRecordUsecase) AddBorrowingRecord(ctx context.Context, br *entity.BorrowingRecords) (*entity.BorrowingRecords, error) {
	atomic := func(c context.Context) error {
		bookCondition := *valueobject.NewCondition("id", valueobject.Equal, br.BookId)
		bookQuery := valueobject.Query{
			Lock:       true,
			Conditions: []valueobject.Condition{bookCondition},
		}
		book := u.bookRepo.First(c, bookQuery)
		if book == nil {
			return apperror.ErrNotFound{
				Resource: "book",
				Field:    "id",
				Value:    br.BookId,
			}
		}
		if book.Quantity == 0 {
			return apperror.ErrEmptyStock{Resource: "book"}
		}
		a, err := u.borrowingRepo.Create(c, br)
		if err != nil {
			return err
		}
		book.Quantity = book.Quantity - 1
		book, err = u.bookRepo.Update(c, book)
		if err != nil {
			return err
		}
		br = a
		return nil
	}
	err := u.borrowingRepo.Run(ctx, atomic)
	if err != nil {
		return nil, err
	}
	return br, nil
}
