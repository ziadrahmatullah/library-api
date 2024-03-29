package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/appjwt"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/hasher"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, user *entity.User) (string, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
	jwt      appjwt.Jwt
	hash     hasher.Hasher
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
	jwt appjwt.Jwt,
	hash hasher.Hasher,
) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		jwt:      jwt,
		hash:     hash,
	}
}

func (u *authUsecase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	emailQuery := valueobject.NewQuery().Condition("email", valueobject.Equal, user.Email)
	fetchedUser, err := u.userRepo.First(ctx, emailQuery)
	if err != nil {
		return nil, err
	}
	if fetchedUser != nil {
		return nil, apperror.NewResourceAlreadyExist("user", "email", user.Email)
	}

	phoneQuery := valueobject.NewQuery().Condition("phone", valueobject.Equal, user.Phone)
	fetchedUser, err = u.userRepo.First(ctx, phoneQuery)
	if err != nil {
		return nil, err
	}
	if fetchedUser != nil {
		return nil, apperror.NewResourceAlreadyExist("user", "phone", user.Phone)
	}

	hashedPassword, err := u.hash.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	createdUser, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil

}

func (u *authUsecase) Login(ctx context.Context, user *entity.User) (string, error) {
	query := valueobject.NewQuery().Condition("email", valueobject.Equal, user.Email)
	fetchedUser, err := u.userRepo.First(ctx, query)
	if err != nil {
		return "", err
	}
	if fetchedUser == nil {
		return "", apperror.NewInvalidCredentialsError()
	}
	if !u.hash.Compare(fetchedUser.Password, user.Password) {
		return "", apperror.NewInvalidCredentialsError()
	}
	token, err := u.jwt.GenerateToken(fetchedUser)
	if err != nil {
		return "", err
	}
	return token, nil
}
