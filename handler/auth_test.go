package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var user = &entity.User{
	Id: 1, Name: "Alice", Email: "alice@example.com", Password: "alice123", Phone: "085234748462",
}

type AuthHandlerTestSuite struct {
	suite.Suite
	router      http.Handler
	rec         *httptest.ResponseRecorder
	authHandler *handler.AuthHandler
	authUsecase *mocks.AuthUsecase
}

func (s *AuthHandlerTestSuite) SetupSubTest() {
	s.authUsecase = mocks.NewAuthUsecase(s.T())
	s.authHandler = handler.NewAuthHandler(s.authUsecase)
	handlers := router.Handlers{
		Auth: s.authHandler,
	}
	s.router = router.New(handlers)
	s.rec = httptest.NewRecorder()
}

func (s *AuthHandlerTestSuite) TestAuthHandler_Register() {
	s.Run("should return 201", func() {
		request := dto.RegisterRequest{
			Email:    user.Email,
			Name:     user.Name,
			Phone:    user.Phone,
			Password: user.Password,
		}
		response := dto.NewFromUser(user)
		responseJson := marshal(gin.H{"data": response})
		s.authUsecase.On("Register", mock.Anything, mock.Anything).Return(user, nil)

		req, _ := http.NewRequest(http.MethodPost, "/register", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusCreated, s.rec.Code)
		s.Equal(responseJson, getBody(s.rec))
	})
	s.Run("should return 400 when not sending required field", func() {
		request := dto.RegisterRequest{
			Email: user.Email,
			Name:  user.Name,
		}

		req, _ := http.NewRequest(http.MethodPost, "/register", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
	s.Run("should return 400 when user with same email or phone already exist", func() {
		request := dto.RegisterRequest{
			Email:    user.Email,
			Name:     user.Name,
			Phone:    user.Phone,
			Password: user.Password,
		}
		s.authUsecase.On("Register", mock.Anything, mock.Anything).Return(nil, apperror.NewResourceAlreadyExist("user", "email", user.Email))

		req, _ := http.NewRequest(http.MethodPost, "/register", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
	s.Run("should return 500 when there's a server error when register new user", func() {
		request := dto.RegisterRequest{
			Email:    user.Email,
			Name:     user.Name,
			Phone:    user.Phone,
			Password: user.Password,
		}
		s.authUsecase.On("Register", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		req, _ := http.NewRequest(http.MethodPost, "/register", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusInternalServerError, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
}

func (s *AuthHandlerTestSuite) TestAuthHandler_Login() {
	s.Run("should return 200", func() {
		request := dto.RegisterRequest{
			Email:    user.Email,
			Password: user.Password,
		}
		response := "token"
		responseJson := marshal(gin.H{"data": response})
		s.authUsecase.On("Login", mock.Anything, mock.Anything).Return("token", nil)

		req, _ := http.NewRequest(http.MethodPost, "/login", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusOK, s.rec.Code)
		s.Equal(responseJson, getBody(s.rec))
	})
	s.Run("should return 400", func() {
		request := dto.RegisterRequest{
			Email: user.Email,
		}

		req, _ := http.NewRequest(http.MethodPost, "/login", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusBadRequest, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
	s.Run("should return 401 when password is wrong", func() {
		request := dto.RegisterRequest{
			Email:    user.Email,
			Password: "wrong password",
		}
		s.authUsecase.On("Login", mock.Anything, mock.Anything).Return("", apperror.NewInvalidCredentialsError())

		req, _ := http.NewRequest(http.MethodPost, "/login", sendBody(request))
		s.router.ServeHTTP(s.rec, req)

		s.Equal(http.StatusUnauthorized, s.rec.Code)
		s.Contains(getBody(s.rec), "error")
	})
}

func TestNewAuthHandler(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}
