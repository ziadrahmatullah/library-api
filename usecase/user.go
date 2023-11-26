package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type UserUsecase interface {
	GetUsers(ctx context.Context, query valueobject.Query) []*entity.User
	GetSingleUser(ctx context.Context, query valueobject.Query) *entity.User
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetUsers(ctx context.Context, query valueobject.Query) []*entity.User {
	return u.userRepo.Find(ctx, query)
}

func (u *userUsecase) GetSingleUser(ctx context.Context, query valueobject.Query) *entity.User {
	return u.userRepo.First(ctx, query)
}
