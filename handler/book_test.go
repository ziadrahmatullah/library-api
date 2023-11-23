package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/dto"
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
		s.bu.On("GetAllBooks", mock.AnythingOfType("valueobject.Clause")).Return([]*entity.Book{})

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
func (s *BookHandlerTestSuite) TestAddBook() {
	s.Run("should return 201", func() {
		quantity := 1
		request := dto.BookRequest{
			Title:       "A",
			Description: "B",
			Quantity:    &quantity,
			AuthorId:    1,
		}
		body, _ := json.Marshal(request)
		s.bu.On("AddBook", mock.AnythingOfType("*entity.Book")).Return(&entity.Book{}, nil)

		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusCreated, s.rec.Code)
	})
	s.Run("should return 400", func() {
		quantity := -1
		request := dto.BookRequest{
			Title:       "A",
			Description: "B",
			Quantity:    &quantity,
			AuthorId:    1,
		}
		body, _ := json.Marshal(request)

		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
	})
	s.Run("should return 409", func() {
		quantity := 1
		request := dto.BookRequest{
			Title:       "A",
			Description: "B",
			Quantity:    &quantity,
			AuthorId:    1,
		}
		body, _ := json.Marshal(request)
		s.bu.On("AddBook", mock.AnythingOfType("*entity.Book")).Return(nil, apperror.ErrAlreadyExist{})

		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusConflict, s.rec.Code)
	})
	s.Run("should return 500", func() {
		quantity := 1
		request := dto.BookRequest{
			Title:       "A",
			Description: "B",
			Quantity:    &quantity,
			AuthorId:    1,
		}
		body, _ := json.Marshal(request)
		s.bu.On("AddBook", mock.AnythingOfType("*entity.Book")).Return(nil, errors.New(""))

		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusInternalServerError, s.rec.Code)
	})
}

func TestBookHandler(t *testing.T) {
	suite.Run(t, new(BookHandlerTestSuite))
}
