package usecase

import (
	"context"
	"database/sql"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BorrowingRecordUsecase interface {
	BorrowBook(ctx context.Context, br *entity.BorrowingRecords) (*entity.BorrowingRecords, error)
	ReturnBook(ctx context.Context, id uint) (*entity.BorrowingRecords, error)
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

func (u *borrowingRecordUsecase) BorrowBook(ctx context.Context, br *entity.BorrowingRecords) (*entity.BorrowingRecords, error) {
	atomic := func(c context.Context) error {
		bookCondition := valueobject.NewCondition("id", valueobject.Equal, br.BookId)
		bookQuery := &valueobject.Query{
			Lock:       true,
			Conditions: []*valueobject.Condition{bookCondition},
		}
		book := u.bookRepo.First(c, bookQuery)
		if book == nil {
			return apperror.Type{
				Type: apperror.NotFound,
				AppError: apperror.ErrNotFound{
					Resource: "book",
					Field:    "id",
					Value:    br.BookId,
				},
			}
		}
		if book.Quantity == 0 {
			return apperror.Type{
				Type:     apperror.BadRequest,
				AppError: apperror.ErrEmptyStock{Resource: "book"},
			}
		}
		br.BorrowedDate = time.Now()
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

func (u *borrowingRecordUsecase) ReturnBook(ctx context.Context, id uint) (*entity.BorrowingRecords, error) {
	var borrowingRecord *entity.BorrowingRecords
	atomic := func(c context.Context) error {
		brCondition := valueobject.NewCondition("id", valueobject.Equal, id)
		brQuery := &valueobject.Query{
			Lock:       true,
			Conditions: []*valueobject.Condition{brCondition},
		}
		br := u.borrowingRepo.First(c, brQuery)
		if br == nil {
			return apperror.Type{
				Type: apperror.NotFound,
				AppError: apperror.ErrNotFound{
					Resource: "borrowing record",
					Field:    "id",
					Value:    id,
				},
			}
		}
		if br.ReturnedDate.Valid {
			return apperror.Type{
				Type:     apperror.Conflict,
				AppError: apperror.ErrBookAlreadyReturned{},
			}
		}
		returnedDate := sql.NullTime{Time: time.Now(), Valid: true}
		br.ReturnedDate = valueobject.NullTime{NullTime: returnedDate}
		br, err := u.borrowingRepo.Update(c, br)
		borrowingRecord = br
		if err != nil {
			return err
		}
		bookCondition := valueobject.NewCondition("id", valueobject.Equal, br.BookId)
		bookQuery := &valueobject.Query{
			Lock:       true,
			Conditions: []*valueobject.Condition{bookCondition},
		}
		book := u.bookRepo.First(c, bookQuery)
		if book == nil {
			return apperror.Type{
				Type: apperror.NotFound,
				AppError: apperror.ErrNotFound{
					Resource: "book",
					Field:    "id",
					Value:    br.BookId,
				},
			}
		}
		book.Quantity = book.Quantity + 1
		book, err = u.bookRepo.Update(c, book)
		if err != nil {
			return err
		}
		return nil
	}
	err := u.borrowingRepo.Run(ctx, atomic)
	if err != nil {
		return nil, err
	}
	return borrowingRecord, nil
}
