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
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
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

var addedBook = models.Book{
	Title:       "Buku 1",
	AuthorId:    1,
	Description: "Tentang Orang 1",
	Quantity:    10,
	Cover:       "Book_One",
	Author:      models.Author{Name: "Ziad"},
}

func removeNewLine(str string) string {
	return strings.Trim(str, "\n")
}

func TestHandleGetBooks(t *testing.T) {
	t.Run("should return 200 if get all books success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: books,
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetAllBooks").Return(books, nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

	t.Run("should return 200 if get all books by title success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: books,
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetBooksByTitle", mock.AnythingOfType("string")).Return(books, nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books?title=buku", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

	t.Run("should return 200 with empty book list", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: make([]models.Book, 0),
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetAllBooks").Return(make([]models.Book, 0), nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

	t.Run("should return 200 if get all books by title success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: make([]models.Book, 0),
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetBooksByTitle", mock.AnythingOfType("string")).Return(make([]models.Book, 0), nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books?title=buku", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetAllBooks").Return(nil, errors.New("Fake error"))
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestHandleCreateBooks(t *testing.T) {
	t.Run("should return 200 if create success", func(t *testing.T) {
		// expectedResp, _ := json.Marshal(dto.Response{
		// 	Data: books,
		// })
		body, _ := json.Marshal(addedBook)
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("CreateBook", mock.IsType(models.Book{})).Return(&addedBook, nil)
		// bookUseCase.On("GetBooksByTitle", mock.AnythingOfType("string")).Return(books, nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodPost, "/books", strings.NewReader(string(body)))
		rec := httptest.NewRecorder()

		r.POST("/books", bookHandler.HandleCreateBook)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

	t.Run("should return 200 if get all books by title success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: books,
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetBooksByTitle", mock.AnythingOfType("string")).Return(books, nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books?title=buku", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

	t.Run("should return 200 with empty book list", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: make([]models.Book, 0),
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetAllBooks").Return(make([]models.Book, 0), nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

	t.Run("should return 200 if get all books by title success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: make([]models.Book, 0),
		})
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("GetBooksByTitle", mock.AnythingOfType("string")).Return(make([]models.Book, 0), nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books?title=buku", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		bookUseCase := mocks.NewBookUsecase(t)
		bookHandler := handler.NewBookHandler(bookUseCase)
		bookUseCase.On("CreateBook", &models.Book{}).Return(nil, nil)
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()

		r.GET("/books", bookHandler.HandleGetBooks)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
