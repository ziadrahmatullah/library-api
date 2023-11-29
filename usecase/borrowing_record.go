package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/transactor"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BorrowingRecordUsecase interface {
	BorrowBook(ctx context.Context, br *entity.BorrowingRecords) (*entity.BorrowingRecords, error)
	ReturnBook(ctx context.Context, id uint) (*entity.BorrowingRecords, error)
}

type borrowingRecordUsecase struct {
	manager       transactor.Manager
	borrowingRepo repository.BorrowingRecordRepository
	bookRepo      repository.BookRepository
}

func NewBorrowingRecordUsecase(manager transactor.Manager, borrowingRepo repository.BorrowingRecordRepository, bookRepo repository.BookRepository) BorrowingRecordUsecase {
	return &borrowingRecordUsecase{
		borrowingRepo: borrowingRepo,
		bookRepo:      bookRepo,
		manager:       manager,
	}
}

func (u *borrowingRecordUsecase) BorrowBook(ctx context.Context, br *entity.BorrowingRecords) (*entity.BorrowingRecords, error) {
	atomic := func(c context.Context) error {
		bookQuery := valueobject.NewQuery().Condition("id", valueobject.Equal, br.BookId).Lock()
		book, err := u.bookRepo.First(c, bookQuery)
		if err != nil {
			return err
		}
		if book == nil {
			return apperror.NewResourceNotFound("book", "id", br.BookId)
		}
		if book.Quantity == 0 {
			return apperror.NewUnavailableResourceError("book")
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
	err := u.manager.Run(ctx, atomic)
	if err != nil {
		return nil, err
	}
	return br, nil
}

func (u *borrowingRecordUsecase) ReturnBook(ctx context.Context, id uint) (*entity.BorrowingRecords, error) {
	var borrowingRecord *entity.BorrowingRecords
	atomic := func(c context.Context) error {
		brQuery := valueobject.NewQuery().Condition("id", valueobject.Equal, id).Lock()
		br, err := u.borrowingRepo.First(c, brQuery)
		if err != nil {
			return err
		}
		if br == nil {
			return apperror.NewResourceNotFound("borrowing record", "id", id)
		}
		if br.ReturnedDate.Valid {
			return fmt.Errorf("book already returned")
		}
		returnedDate := sql.NullTime{Time: time.Now(), Valid: true}
		br.ReturnedDate = valueobject.NullTime{NullTime: returnedDate}
		br, err = u.borrowingRepo.Update(c, br)
		if err != nil {
			return err
		}
		borrowingRecord = br

		bookQuery := valueobject.NewQuery().Condition("id", valueobject.Equal, br.BookId).Lock()
		book, err := u.bookRepo.First(c, bookQuery)
		if err != nil {
			return err
		}
		if book == nil {
			return apperror.NewResourceNotFound("book", "id", br.BookId)
		}
		book.Quantity = book.Quantity + 1
		book, err = u.bookRepo.Update(c, book)
		if err != nil {
			return err
		}
		return nil
	}
	err := u.manager.Run(ctx, atomic)
	if err != nil {
		return nil, err
	}
	return borrowingRecord, nil
}
