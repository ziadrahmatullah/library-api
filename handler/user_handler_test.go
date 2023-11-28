package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

var users = []models.User{
	{
		Name: "Alice",
		Email: "alice@gmail.com",
		Phone: "0823728327",
	},
}

func TestHandleGetUsers(t *testing.T){
	t.Run("should return 200 if get all users success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: users,
		})
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewUserHandler(uu)
		rec := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/users", nil)
		uu.On("GetAllUsers", c).Return(users, nil)

		uh.HandleGetUsers(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 200 if get all users by name success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: users,
		})
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewUserHandler(uu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/users?name=Alice", nil)
		uu.On("GetUserByName", c, "Alice").Return(users, nil)

		uh.HandleGetUsers(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 200 with empty book list", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: make([]models.User, 0),
		})
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewUserHandler(uu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/users", nil)
		uu.On("GetAllUsers", c).Return(make([]models.User, 0), nil)

		uh.HandleGetUsers(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		uu := mocks.NewUserUsecase(t)
		uh := handler.NewUserHandler(uu)
		uu.On("GetAllUsers", mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			UserHandler: uh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleUserRegister(t *testing.T) {
		
}