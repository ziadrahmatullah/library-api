package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/util"
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

var quantity = 10
var negativeQty = -10

var bookReq = dto.BookReq{
	Title:       "Buku1",
	Description: "Tentang orang 1",
	Quantity:    &quantity,
	Cover:       "halo",
	AuthorId:    1,
}

var bookNoCoverReq = dto.BookReq{
	Title:       "Buku1",
	Description: "Tentang orang 1",
	Quantity:    &quantity,
	AuthorId:    1,
}

var invalidBook = dto.BookReq{
	Title:       "Buku1",
	Description: "Tentang orang 1",
	Quantity:    &negativeQty,
	AuthorId:    1,
}

var book = models.Book{
	Title:       "Buku1",
	Description: "Tentang orang 1",
	Quantity:    10,
	Cover:       "halo",
	AuthorId:    1,
}

var bookNoCover = models.Book{
	Title:       "Buku1",
	Description: "Tentang orang 1",
	Quantity:    10,
	Cover:       "",
	AuthorId:    1,
}

func TestHandleGetBooks(t *testing.T) {
	t.Run("should return 200 if get all books success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: books,
		})
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		rec := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/books", nil)
		bu.On("GetAllBooks", c).Return(books, nil)

		bh.HandleGetBooks(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 200 if get all books by title success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: books,
		})
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/books?title=Buku1", nil)
		bu.On("GetBooksByTitle", c, "Buku1").Return(books, nil)

		bh.HandleGetBooks(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 200 with empty book list", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: make([]models.Book, 0),
		})
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/books", nil)
		bu.On("GetAllBooks", c).Return(make([]models.Book, 0), nil)

		bh.HandleGetBooks(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		bu.On("GetAllBooks", mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			BookHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleCreateBooks(t *testing.T) {
	t.Run("should return 200 if create success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: book,
		})
		param, _ := json.Marshal(bookReq)
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		rec := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/books", strings.NewReader(string(param)))
		bu.On("CreateBook", c, bookReq.ToBookModel()).Return(&book, nil)

		bh.HandleCreateBook(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 200 if cover not given", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: bookNoCover,
		})
		param, _ := json.Marshal(bookNoCoverReq)
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		rec := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/books", strings.NewReader(string(param)))
		bu.On("CreateBook", c, bookNoCoverReq.ToBookModel()).Return(&bookNoCover, nil)

		bh.HandleCreateBook(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when title already exist", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "title already exist")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(bookReq)
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		bu.On("CreateBook", mock.Anything, bookReq.ToBookModel()).Return(nil, expectedErr)
		opts := server.RouterOpts{
			BookHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/books", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invalidBook)
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		opts := server.RouterOpts{
			BookHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/books", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(bookReq)
		bu := mocks.NewBookUsecase(t)
		bh := handler.NewBookHandler(bu)
		bu.On("CreateBook", mock.Anything, bookReq.ToBookModel()).Return(nil, expectedErr)
		opts := server.RouterOpts{
			BookHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/books", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}
