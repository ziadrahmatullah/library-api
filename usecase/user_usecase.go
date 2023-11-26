package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]models.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(u repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: u,
	}
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.userRepository.FindUsers()
}
