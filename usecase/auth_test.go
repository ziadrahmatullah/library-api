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
	hash        *mocks.Hasher
	authUsecase usecase.AuthUsecase
}

func (s *AuthUsecaseTestSuite) SetupSubTest() {
	s.userRepo = mocks.NewUserRepository(s.T())
	s.jwt = mocks.NewJwt(s.T())
	s.hash = mocks.NewHasher(s.T())
	s.authUsecase = usecase.NewAuthUsecase(s.userRepo, s.jwt, s.hash)
}

func (s *AuthUsecaseTestSuite) TestAuthUsecase_Register() {
	s.Run("should return new user", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.hash.On("Hash", mock.Anything).Return("", nil)
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
	s.Run("should return error when hashing the password", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.hash.On("Hash", mock.Anything).Return("", errors.New(""))

		createdUser, err := s.authUsecase.Register(context.Background(), user)

		s.Nil(createdUser)
		s.Error(err)
	})
	s.Run("should return error when there's an error when creating new user", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.hash.On("Hash", mock.Anything).Return("", nil)
		s.userRepo.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		ctx := context.WithValue(context.Background(), "", "")

		createdUser, err := s.authUsecase.Register(ctx, user)

		s.Nil(createdUser)
		s.Error(err)
	})
}

func (s *AuthUsecaseTestSuite) TestAuthUsecase_Login() {
	s.Run("should return token", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(user, nil)
		s.hash.On("Compare", mock.Anything, mock.Anything).Return(true)
		s.jwt.On("GenerateToken", mock.Anything).Return("", nil)

		token, err := s.authUsecase.Login(context.Background(), user)

		s.NotNil(token)
		s.NoError(err)
	})
	s.Run("should return error when there's an error when searching for user", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		token, err := s.authUsecase.Login(context.Background(), user)

		s.Equal("", token)
		s.Error(err)
	})
	s.Run("should return error when there's no user", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)

		token, err := s.authUsecase.Login(context.Background(), user)

		s.Equal("", token)
		s.Error(err)
	})
	s.Run("should return error when password is wrong", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(user, nil)
		s.hash.On("Compare", mock.Anything, mock.Anything).Return(false)

		token, err := s.authUsecase.Login(context.Background(), user)

		s.Equal("", token)
		s.Error(err)
	})
	s.Run("should return error when there's an error when generating token", func() {
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(user, nil)
		s.hash.On("Compare", mock.Anything, mock.Anything).Return(true)
		s.jwt.On("GenerateToken", mock.Anything).Return("", errors.New(""))

		token, err := s.authUsecase.Login(context.Background(), user)

		s.Equal("", token)
		s.Error(err)
	})
}

func TestAuthUsecase(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTestSuite))
}
