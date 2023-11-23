package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/router"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BookHandlerTestSuite struct {
	suite.Suite
	router http.Handler
	bu     *mocks.BookUsecase
	bh     *handler.BookHandler
	rec    *httptest.ResponseRecorder
}

func (s *BookHandlerTestSuite) SetupSubTest() {
	s.bu = mocks.NewBookUsecase(s.T())
	s.bh = handler.NewBookHandler(s.bu)
	handlers := router.Handlers{
		Book: s.bh,
	}
	s.router = router.New(handlers)
	s.rec = httptest.NewRecorder()
}

func (s *BookHandlerTestSuite) TestListBooks() {
	s.Run("should return 200", func() {
		s.bu.On("GetAllBooks").Return([]*entity.Book{})

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusOK, s.rec.Code)
	})
	s.Run("should return 200 when search by name", func() {
		s.bu.On("FindBooksByTitle", mock.AnythingOfType("string")).Return([]*entity.Book{})

		req, _ := http.NewRequest(http.MethodGet, "/books?title=how", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusOK, s.rec.Code)
	})
}

func TestBookHandler(t *testing.T) {
	suite.Run(t, new(BookHandlerTestSuite))
}
