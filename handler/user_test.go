package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/router"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var users = []*entity.User{
	{Id: 1, Name: "Alice", Email: "alice@example.com", Phone: "082356783412"},
}

type UserHandlerTestSuite struct {
	suite.Suite
	router http.Handler
	uu     *mocks.UserUsecase
	uh     *handler.UserHandler
	rec    *httptest.ResponseRecorder
}

func (s *UserHandlerTestSuite) SetupSubTest() {
	s.uu = mocks.NewUserUsecase(s.T())
	s.uh = handler.NewUserHandler(s.uu)
	handlers := router.Handlers{
		User: s.uh,
	}
	s.router = router.New(handlers)
	s.rec = httptest.NewRecorder()
}

func (s *UserHandlerTestSuite) TestListUser() {
	s.Run("should return 200", func() {
		s.uu.On("GetUsers", mock.AnythingOfType("valueobject.Query")).Return(users)
		response := h{"data": dto.NewFromUsers(users)}
		responseJson := marshal(response)

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusOK, s.rec.Code)
		s.Equal(responseJson, getBody(s.rec))
	})
	s.Run("should return 400", func() {
		req, _ := http.NewRequest(http.MethodGet, "/users?page=a", nil)
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
