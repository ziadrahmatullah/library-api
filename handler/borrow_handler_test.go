package handler_test

import (
	"encoding/json"
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

var Borrows = []models.BorrowBook{
	{
		UserId:        1,
		BookId:        1,
		Status:        "not complete",
	},
}

var BorrowsCreate = []dto.BorrowReq{
	{
		UserId: 1,
		BookId: 1,
	},
}

var BorrowsResponse = []dto.BorrowRes{
	{
		UserId:        1,
		BookId:        1,
		Status:        "not complete",
	},
}

func TestHandlBorrowBook(t *testing.T) {
	t.Run("should return 200 with borrows when request valid", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: BorrowsResponse[0],
		})
		param, _ := json.Marshal(BorrowsCreate[0])
		bu := mocks.NewBorrowUsecase(t)
		bh := handler.NewBorrowHandler(bu)
		bu.On("BorrowBook", BorrowsCreate[0].ToBorrowModel()).Return(&Borrows[0], nil)
		opts := server.RouterOpts{
			BorrowHandler: bh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/borrows", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})
}