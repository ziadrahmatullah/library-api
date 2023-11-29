package usecase_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	userRepo    *mocks.UserRepository
	userUsecase usecase.UserUsecase
}

func (s *UserUsecaseTestSuite) SetupSubTest() {
	s.userRepo = mocks.NewUserRepository(s.T())
	s.userUsecase = usecase.NewUserUsecase(s.userRepo)
}

func (s *UserUsecaseTestSuite) TestUserUsecase_GetUsers() {
	s.Run("should return users", func() {
		query := valueobject.NewQuery()
		s.userRepo.On("Find", mock.Anything, mock.Anything).Return(users, nil)

		fetchedUsers, err := s.userUsecase.GetUsers(context.Background(), query)

		s.Equal(users, fetchedUsers)
		s.NoError(err)
	})
	s.Run("should return user", func() {
		query := valueobject.NewQuery()
		s.userRepo.On("First", mock.Anything, mock.Anything).Return(user, nil)

		fetchedUser, err := s.userUsecase.GetSingleUser(context.Background(), query)

		s.Equal(user, fetchedUser)
		s.NoError(err)
	})
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
