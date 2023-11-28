package handler_test

// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/mocks"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/server"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/util"
// 	"github.com/go-playground/assert/v2"
// )

// var borrows = []models.BorrowBook{
// 	{
// 		UserId: 1,
// 		BookId: 1,
// 		Status: "not returned",
// 	},
// }

// var borrowsReq = []dto.BorrowReq{
// 	{
// 		UserId: 1,
// 		BookId: 1,
// 	},
// }

// var borrowsRes = []dto.BorrowRes{
// 	{
// 		UserId: 1,
// 		BookId: 1,
// 		Status: "not returned",
// 	},
// }

// func TestHandleGetRecords(t *testing.T) {
// 	t.Run("should return 200 when get records success", func(t *testing.T) {
// 		expectedResp, _ := json.Marshal(dto.Response{
// 			Data: borrows,
// 		})
// 		borrowUsecase := mocks.NewBorrowUsecase(t)
// 		borrowHandler := handler.NewBorrowHandler(borrowUsecase)
// 		borrowUsecase.On("GetAllRecords").Return(borrows, nil)
// 		opts := server.RouterOpts{
// 			BorrowHandler: borrowHandler,
// 		}
// 		r := server.NewRouter(opts)

// 		req, _ := http.NewRequest(http.MethodGet, "/borrows", nil)
// 		rec := httptest.NewRecorder()
// 		r.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
// 	})

// 	t.Run("should return 500 while error in server", func(t *testing.T) {
// 		borrowUsecase := mocks.NewBorrowUsecase(t)
// 		borrowHandler := handler.NewBorrowHandler(borrowUsecase)
// 		borrowUsecase.On("GetAllRecords").Return(nil, errors.New("Fake error"))
// 		opts := server.RouterOpts{
// 			BorrowHandler: borrowHandler,
// 		}
// 		r := server.NewRouter(opts)

// 		req, _ := http.NewRequest(http.MethodGet, "/borrows", nil)
// 		rec := httptest.NewRecorder()
// 		r.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 	})
// }

// func TestHandlBorrowBook(t *testing.T) {
// 	t.Run("should return 200 when borrow book success", func(t *testing.T) {
// 		expectedResp, _ := json.Marshal(dto.Response{
// 			Data: borrowsRes[0],
// 		})
// 		param, _ := json.Marshal(borrowsReq[0])
// 		borrowUsecase := mocks.NewBorrowUsecase(t)
// 		borrowHandler := handler.NewBorrowHandler(borrowUsecase)
// 		borrowUsecase.On("BorrowBook", borrowsReq[0].ToBorrowModel()).Return(&borrows[0], nil)
// 		opts := server.RouterOpts{
// 			BorrowHandler: borrowHandler,
// 		}
// 		r := server.NewRouter(opts)
// 		rec := httptest.NewRecorder()

// 		req, _ := http.NewRequest(http.MethodPost, "/borrows", strings.NewReader(string(param)))
// 		r.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
// 	})

// 	t.Run("should return 500 while error in server", func(t *testing.T) {
// 		param, _ := json.Marshal(borrowsReq[0])
// 		borrowUsecase := mocks.NewBorrowUsecase(t)
// 		borrowHandler := handler.NewBorrowHandler(borrowUsecase)
// 		borrowUsecase.On("BorrowBook", borrowsReq[0].ToBorrowModel()).Return(nil, errors.New("Fake error"))
// 		opts := server.RouterOpts{
// 			BorrowHandler: borrowHandler,
// 		}
// 		r := server.NewRouter(opts)

// 		req, _ := http.NewRequest(http.MethodPost, "/borrows", strings.NewReader(string(param)))
// 		rec := httptest.NewRecorder()
// 		r.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 	})
// }

// func TestHandlReturnBook(t *testing.T) {
// 	t.Run("should return 200 when return book success", func(t *testing.T) {
// 		expectedResp, _ := json.Marshal(dto.Response{
// 			Data: borrowsRes[0],
// 		})
// 		param, _ := json.Marshal(borrowsReq[0])
// 		borrowUsecase := mocks.NewBorrowUsecase(t)
// 		borrowHandler := handler.NewBorrowHandler(borrowUsecase)
// 		borrowUsecase.On("ReturnBook", borrowsReq[0].ToBorrowModel()).Return(&borrows[0], nil)
// 		opts := server.RouterOpts{
// 			BorrowHandler: borrowHandler,
// 		}
// 		r := server.NewRouter(opts)
// 		rec := httptest.NewRecorder()

// 		req, _ := http.NewRequest(http.MethodPut, "/borrows", strings.NewReader(string(param)))
// 		r.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
// 	})

// 	t.Run("should return 500 while error in server", func(t *testing.T) {
// 		param, _ := json.Marshal(borrowsReq[0])
// 		borrowUsecase := mocks.NewBorrowUsecase(t)
// 		borrowHandler := handler.NewBorrowHandler(borrowUsecase)
// 		borrowUsecase.On("ReturnBook", borrowsReq[0].ToBorrowModel()).Return(nil, errors.New("Fake error"))
// 		opts := server.RouterOpts{
// 			BorrowHandler: borrowHandler,
// 		}
// 		r := server.NewRouter(opts)

// 		req, _ := http.NewRequest(http.MethodPut, "/borrows", strings.NewReader(string(param)))
// 		rec := httptest.NewRecorder()
// 		r.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 	})
// }
