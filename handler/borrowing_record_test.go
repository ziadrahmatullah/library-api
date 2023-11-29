package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/router"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var records = []*entity.BorrowingRecords{
	{Id: 1, BookId: 1, UserId: 1},
}

type BorrowingRecordHandlerTestSuite struct {
	suite.Suite
	router http.Handler
	bru    *mocks.BorrowingRecordUsecase
	brh    *handler.BorrowingRecordHandler
	rec    *httptest.ResponseRecorder
}

func (s *BorrowingRecordHandlerTestSuite) SetupSubTest() {
	s.bru = mocks.NewBorrowingRecordUsecase(s.T())
	s.brh = handler.NewBorrowingRecordHandler(s.bru)
	handlers := router.Handlers{
		BorrowingRecord: s.brh,
	}
	s.router = router.New(handlers)
	s.rec = httptest.NewRecorder()
}

func (s *BorrowingRecordHandlerTestSuite) TestBorrowingRecordHandler_AddBorrowing() {
	s.Run("should return 200", func() {
		request := dto.BorrowingRecordRequest{
			UserId: 1,
			BookId: 1,
		}
		s.bru.On("BorrowBook", mock.Anything, mock.Anything).Return(records[0], nil)
		response := dto.NewFromBorrowingRecord(records[0])
		responseJson := marshal(h{"data": response})

		req, _ := http.NewRequest(http.MethodPost, "/borrowing-records", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusCreated, s.rec.Code)
		s.Equal(responseJson, getBody(s.rec))
	})

	s.Run("should return 400", func() {
		request := dto.BorrowingRecordRequest{
			UserId: 1,
		}

		req, _ := http.NewRequest(http.MethodPost, "/borrowing-records", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})

	s.Run("should return 500", func() {
		request := dto.BorrowingRecordRequest{
			UserId: 1,
			BookId: 1,
		}
		s.bru.On("BorrowBook", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		req, _ := http.NewRequest(http.MethodPost, "/borrowing-records", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusInternalServerError, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
}

func (s *BorrowingRecordHandlerTestSuite) TestBorrowingRecordHandler_ReturnBook() {
	s.Run("should return 200", func() {
		s.bru.On("ReturnBook", mock.Anything, mock.Anything).Return(records[0], nil)
		response := dto.NewFromBorrowingRecord(records[0])
		responseJson := marshal(h{"data": response})

		req, _ := http.NewRequest(http.MethodPut, "/borrowing-records/1", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusOK, s.rec.Code)
		s.Equal(responseJson, getBody(s.rec))
	})
	s.Run("should return 400", func() {
		req, _ := http.NewRequest(http.MethodPut, "/borrowing-records/a", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
	s.Run("should return 500", func() {
		s.bru.On("ReturnBook", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		req, _ := http.NewRequest(http.MethodPut, "/borrowing-records/1", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusInternalServerError, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
}

func TestBorrowingRecordHandler(t *testing.T) {
	suite.Run(t, new(BorrowingRecordHandlerTestSuite))
}
