package usecase_test

import (
	"context"
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var users = []*entity.User{
	{Id: 1, Name: "Alice", Email: "alice@example.com", Password: "alice123", Phone: "08534567822"},
}

var user = users[0]

type AuthUsecaseTestSuite struct {
	suite.Suite
	userRepo    *mocks.UserRepository
	jwt         *mocks.Jwt
	authUsecase usecase.AuthUsecase
}

func (s *AuthUsecaseTestSuite) SetupSubTest() {
	s.userRepo = mocks.NewUserRepository(s.T())
	s.jwt = mocks.NewJwt(s.T())
	s.authUsecase = usecase.NewAuthUsecase(s.userRepo, s.jwt)
}

func (s *AuthUsecaseTestSuite) TestAuthUsecase_Login() {
	s.Run("should return new user", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.userRepo.On("Create", mock.Anything, mock.Anything).Return(user, nil)

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.authUsecase.Register(ctx, user)

		s.Equal(user, createdUser)
		s.NoError(err)
	})
	s.Run("should return error when there's an error when fetching user by email", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.authUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
	s.Run("should return error when user with same email already registered", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(user, nil)

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.authUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
	s.Run("should return error when there's an error when fetching user by phone", func() {
		emailQuery := valueobject.NewQuery().Condition("email", valueobject.Equal, user.Email)
		phoneQuery := valueobject.NewQuery().Condition("phone", valueobject.Equal, user.Phone)
		s.userRepo.On("First", mock.Anything, emailQuery).Return(nil, nil)
		s.userRepo.On("First", mock.Anything, phoneQuery).Return(nil, errors.New(""))

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.authUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
	s.Run("should return error when user with same phone already registered", func() {
		emailQuery := valueobject.NewQuery().Condition("email", valueobject.Equal, user.Email)
		phoneQuery := valueobject.NewQuery().Condition("phone", valueobject.Equal, user.Phone)
		s.userRepo.On("First", mock.Anything, emailQuery).Return(nil, nil)
		s.userRepo.On("First", mock.Anything, phoneQuery).Return(user, nil)

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.authUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
	s.Run("should return error when there's an error when creating new user", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.userRepo.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.authUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
}

func TestAuthUsecase(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTestSuite))
}
