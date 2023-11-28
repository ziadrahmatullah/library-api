package handler_test

// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/mocks"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/server"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/util"
// 	"github.com/go-playground/assert/v2"
// )

// var users = []models.User{
// 	{
// 		Name: "Alice",
// 		Email: "alice@gmail.com",
// 		Phone: "0823728327",
// 	},
// }

// func TestHandleGetUsers(t *testing.T){
// 	t.Run("should return 200 when get records success", func(t *testing.T) {
// 		expectedResp, _ := json.Marshal(dto.Response{
// 			Data: users,
// 		})
// 		userUsecase := mocks.NewUserUsecase(t)
// 		userHandler := handler.NewUserHandler(userUsecase)
// 		userUsecase.On("GetAllUsers").Return(users, nil)
// 		opts := server.RouterOpts{
// 			UserHandler: userHandler,
// 		}
// 		r := server.NewRouter(opts)
		
// 		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
// 		rec := httptest.NewRecorder()
// 		r.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
// 	})

// 	t.Run("should return 500 while error in server", func(t *testing.T) {
// 		userUsecase := mocks.NewUserUsecase(t)
// 		userHandler := handler.NewUserHandler(userUsecase)
// 		userUsecase.On("GetAllUsers").Return(nil, errors.New("Fake error"))
// 		opts := server.RouterOpts{
// 			UserHandler: userHandler,
// 		}
// 		r := server.NewRouter(opts)
		
// 		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
// 		rec := httptest.NewRecorder()
// 		r.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 	})
// }