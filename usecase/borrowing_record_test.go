package usecase_test

import (
	"context"
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var borrowingRecord = &entity.BorrowingRecords{
	Id:     1,
	UserId: 1,
	BookId: 1,
}

type BorrowingRecordUsecaseTestSuite struct {
	suite.Suite
	bru     usecase.BorrowingRecordUsecase
	brr     *mocks.BorrowingRecordRepository
	br      *mocks.BookRepository
	manager *mocks.Manager
}

func (s *BorrowingRecordUsecaseTestSuite) SetupSubTest() {
	s.manager = mocks.NewManager(s.T())
	s.brr = mocks.NewBorrowingRecordRepository(s.T())
	s.br = mocks.NewBookRepository(s.T())
	s.bru = usecase.NewBorrowingRecordUsecase(s.manager, s.brr, s.br)
}

func (s *BorrowingRecordUsecaseTestSuite) TestBorrowingRecordUsecase_BorrowBook() {
	s.Run("should return borrowing record", func() {
		s.manager.On("Run", mock.Anything, mock.Anything).Return(nil)

		createdBr, err := s.bru.BorrowBook(context.Background(), borrowingRecord)

		s.Equal(borrowingRecord, createdBr)
		s.NoError(err)
	})
	s.Run("should return error when transaction fail", func() {
		s.manager.On("Run", mock.Anything, mock.Anything).Return(errors.New(""))

		createdBr, err := s.bru.BorrowBook(context.Background(), borrowingRecord)

		s.Nil(createdBr)
		s.Error(err)
	})
}

func (s *BorrowingRecordUsecaseTestSuite) TestBorrowingRecordUsecase_ReturnBook() {
	s.Run("should return borrowing record", func() {
		s.manager.On("Run", mock.Anything, mock.Anything).Return(nil)

		_, err := s.bru.ReturnBook(context.Background(), borrowingRecord.Id)

		s.NoError(err)
	})
	s.Run("should return error when transaction fail", func() {
		s.manager.On("Run", mock.Anything, mock.Anything).Return(errors.New(""))

		createdBr, err := s.bru.ReturnBook(context.Background(), borrowingRecord.Id)

		s.Nil(createdBr)
		s.Error(err)
	})
}

func TestBorrowingRecordUsecase(t *testing.T) {
	suite.Run(t, new(BorrowingRecordUsecaseTestSuite))
}
