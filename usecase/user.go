package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type UserUsecase interface {
	GetUsers(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.User
	GetSingleUser(conditions []valueobject.Condition) *entity.User
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetUsers(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.User {
	return u.userRepo.Find(clause, conditions)
}

func (u *userUsecase) GetSingleUser(conditions []valueobject.Condition) *entity.User {
	return u.userRepo.First(conditions)
}
