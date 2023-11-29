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

var borrowsReq = []dto.BorrowReq{
	{
		BookId: 1,
	},
	{
	},
}

var borrowsRes = []dto.BorrowRes{
	{
		UserId: 1,
		BookId: 1,
		Status: "not returned",
	},
}

var borrowRecord = []models.BorrowBook{
	{
		UserId: 1,
		BookId: 1,
		Status: "not returned",
	},
}

func TestHandleGetRecords(t *testing.T) {
	t.Run("should return 200 if get all records success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: borrowRecord,
		})
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/borrows", nil)
		bu.On("GetAllRecords", c).Return(borrowRecord, nil)

		bh.HandleGetRecords(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		bu.On("GetAllRecords", mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			BorrowHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/borrows", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleBorrowBook(t *testing.T) {
	t.Run("should return 200 if borrow success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: borrowsRes[0],
		})
		param, _ := json.Marshal(borrowsReq[0])
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/borrows", strings.NewReader(string(param)))
		bu.On("BorrowBook", c, borrowsReq[0].ToBorrowModel(0)).Return(&borrowRecord[0], nil)

		bh.HandleBorrowBook(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(borrowsReq[1])
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		opts := server.RouterOpts{
			BorrowHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/borrows", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(borrowsReq[0])
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		bu.On("BorrowBook", mock.Anything, borrowsReq[0].ToBorrowModel(0)).Return(nil, expectedErr)
		opts := server.RouterOpts{
			BorrowHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/borrows", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleReturnBook(t *testing.T) {
	t.Run("should return 200 if borrow success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: borrowsRes[0],
		})
		param, _ := json.Marshal(borrowsReq[0])
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPut, "/borrows", strings.NewReader(string(param)))
		bu.On("ReturnBook", c, borrowsReq[0].ToBorrowModel(0)).Return(&borrowRecord[0], nil)

		bh.HandleReturnBook(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(borrowsReq[1])
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		opts := server.RouterOpts{
			BorrowHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPut, "/borrows", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(borrowsReq[0])
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		bu.On("ReturnBook", mock.Anything, borrowsReq[0].ToBorrowModel(0)).Return(nil, expectedErr)
		opts := server.RouterOpts{
			BorrowHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPut, "/borrows", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}
