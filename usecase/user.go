package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type UserUsecase interface {
	GetUsers(query valueobject.Query) []*entity.User
	GetSingleUser(query valueobject.Query) *entity.User
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetUsers(query valueobject.Query) []*entity.User {
	return u.userRepo.Find(query)
}

func (u *userUsecase) GetSingleUser(query valueobject.Query) *entity.User {
	return u.userRepo.First(query)
}
