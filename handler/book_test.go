package handler_test

import (
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

var books = []*entity.Book{
	{Id: 1, Title: "A", Description: "B", Quantity: 1},
}

type BookHandlerTestSuite struct {
	suite.Suite
	router http.Handler
	bu     *mocks.BookUsecase
	bh     *handler.BookHandler
	rec    *httptest.ResponseRecorder
}

var emptyCtx = mock.AnythingOfType("*context.emptyCtx")

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
		s.bu.On("GetAllBooks", emptyCtx, mock.AnythingOfType("valueobject.Query")).Return(books)
		response := dto.NewFromBooks(books)
		responseJson := marshal(h{"data": response})

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusOK, s.rec.Code)
		s.Equal(responseJson, getBody(s.rec))
	})
	s.Run("should return 200 when search by name", func() {
		s.bu.On("GetAllBooks", emptyCtx, mock.AnythingOfType("valueobject.Query")).Return([]*entity.Book{})
		response := make([]*dto.BookResponse, 0)
		responseJson := marshal(h{"data": response})

		req, _ := http.NewRequest(http.MethodGet, "/books?title=how", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusOK, s.rec.Code)
		s.Equal(responseJson, getBody(s.rec))
	})
	s.Run("should return 400 when param is invalid", func() {
		req, _ := http.NewRequest(http.MethodGet, "/books?page=a", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
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
		s.bu.On("AddBook", emptyCtx, mock.AnythingOfType("*entity.Book")).Return(books[0], nil)
		response := dto.NewFromBook(books[0])
		responseJson := marshal(h{"data": response})

		req, _ := http.NewRequest(http.MethodPost, "/books", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusCreated, s.rec.Code)
		s.Equal(responseJson, getBody(s.rec))
	})
	s.Run("should return 400", func() {
		quantity := -1
		request := dto.BookRequest{
			Title:       "A",
			Description: "B",
			Quantity:    &quantity,
			AuthorId:    1,
		}

		req, _ := http.NewRequest(http.MethodPost, "/books", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
	s.Run("should return 409", func() {
		quantity := 1
		request := dto.BookRequest{
			Title:       "A",
			Description: "B",
			Quantity:    &quantity,
			AuthorId:    1,
		}
		s.bu.On("AddBook", emptyCtx, mock.AnythingOfType("*entity.Book")).Return(nil, apperror.Type{Type: apperror.Conflict, AppError: apperror.ErrAlreadyExist{}})

		req, _ := http.NewRequest(http.MethodPost, "/books", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusConflict, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
	s.Run("should return 500", func() {
		quantity := 1
		request := dto.BookRequest{
			Title:       "A",
			Description: "B",
			Quantity:    &quantity,
			AuthorId:    1,
		}
		s.bu.On("AddBook", emptyCtx, mock.AnythingOfType("*entity.Book")).Return(nil, errors.New(""))

		req, _ := http.NewRequest(http.MethodPost, "/books", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusInternalServerError, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
}

func TestBookHandler(t *testing.T) {
	suite.Run(t, new(BookHandlerTestSuite))
}
