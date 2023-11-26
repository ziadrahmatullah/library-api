package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/util"
	"github.com/go-playground/assert/v2"
)

var books = []models.Book{
	{
		Title:       "Buku 1",
		Description: "Tentang orang 1",
		Quantity:    10,
		Cover:       "",
		AuthorId:    1,
	},
	{
		Title:       "Buku 2",
		Description: "Tentang orang 2",
		Quantity:    10,
		Cover:       "",
		AuthorId:    1,
	},
}

var quantity = 10

var bookReq = dto.BookReq{
	Title:       "Buku1",
	Description: "Tentang orang 1",
	Quantity:    &quantity,
	Cover:       "halo",
	AuthorId:    1,
}

var book = models.Book{
	Title:       "Buku1",
	Description: "Tentang orang 1",
	Quantity:    10,
	Cover:       "halo",
	AuthorId:    1,
}


func TestHandleGetBooks(t *testing.T) {
	t.Run("should return 200 if get all books success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: books,
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetAllBooks").Return(books, nil)
		opts := server.RouterOpts{
			BookHandler: bookHandler,
		}
		r := server.NewRouter(opts)
		
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 200 if get all books by title success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: books,
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetBooksByTitle", "Buku1").Return(books, nil)
		opts := server.RouterOpts{
			BookHandler: bookHandler,
		}
		r := server.NewRouter(opts)
		
		req, _ := http.NewRequest(http.MethodGet, "/books?title=Buku1", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 200 with empty book list", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: make([]models.Book, 0),
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetAllBooks").Return(make([]models.Book, 0), nil)
		opts := server.RouterOpts{
			BookHandler: bookHandler,
		}
		r := server.NewRouter(opts)
		
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetAllBooks").Return(nil, errors.New("Fake error"))
		opts := server.RouterOpts{
			BookHandler: bookHandler,
		}
		r := server.NewRouter(opts)

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestHandleCreateBooks(t *testing.T) {
	t.Run("should return 200 if create success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: book,
		})
		body, _ := json.Marshal(bookReq)
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("CreateBook", book).Return(&book, nil)
		opts := server.RouterOpts{
			BookHandler: bookHandler,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()
		
		req, _ := http.NewRequest(http.MethodPost, "/books", strings.NewReader(string(body)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		body, _ := json.Marshal(bookReq)
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("CreateBook", book).Return(nil, errors.New("Fake error"))
		opts := server.RouterOpts{
			BookHandler: bookHandler,
		}
		r := server.NewRouter(opts)

		req, _ := http.NewRequest(http.MethodPost, "/books", strings.NewReader(string(body)))
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
