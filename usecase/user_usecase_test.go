package usecase_test

import (
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var users = []models.User{
	{
		Name:  "Alice",
		Email: "alice@gmail.com",
		Phone: "0823728327",
	},
}

func TestGetAllUsers(t *testing.T) {
	t.Run("should return users when success", func(t *testing.T) {
		userRepository := mocks.NewUserRepository(t)
		userUsecase := usecase.NewUserUsecase(userRepository)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		userRepository.On("FindUsers", c).Return(users, nil)

		resUsers, _ := userUsecase.GetAllUsers(c)

		assert.Equal(t, users, resUsers)
	})
}

func TestGetUserByName(t *testing.T) {
	t.Run("should return users when success", func(t *testing.T) {
		userRepository := mocks.NewUserRepository(t)
		userUsecase := usecase.NewUserUsecase(userRepository)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		userRepository.On("FindUserByName", c, "Alice").Return(users, nil)

		resUsers, _ := userUsecase.GetUserByName(c, "Alice")

		assert.Equal(t, users, resUsers)
	})
}
