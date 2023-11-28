package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/appjwt"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, user *entity.User) (string, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
	jwt      appjwt.Jwt
}

func NewAuthUsecase(userRepo repository.UserRepository, jwt appjwt.Jwt) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (u *authUsecase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	emailQuery := valueobject.NewQuery().Condition("email", valueobject.Equal, user.Email)
	fetchedUser, err := u.userRepo.First(ctx, emailQuery)
	if err != nil {
		return nil, err
	}
	if fetchedUser != nil {
		return nil, apperror.Type{
			Type: apperror.Conflict,
			AppError: apperror.ErrAlreadyExist{
				Resource: "user",
				Field:    "email",
				Value:    user.Email,
			},
		}
	}

	phoneQuery := valueobject.NewQuery().Condition("phone", valueobject.Equal, user.Phone)
	fetchedUser, err = u.userRepo.First(ctx, phoneQuery)
	if err != nil {
		return nil, err
	}
	if fetchedUser != nil {
		return nil, apperror.Type{
			Type: apperror.Conflict,
			AppError: apperror.ErrAlreadyExist{
				Resource: "user",
				Field:    "phone",
				Value:    user.Phone,
			},
		}
	}

	hashedPassword, err := hashPassword(user.Password)
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
		return "", apperror.Type{
			Type:     apperror.UnAuthenticated,
			AppError: apperror.ErrInvalidCredential{},
		}
	}
	if !checkPasswordHash(fetchedUser.Password, user.Password) {
		return "", apperror.Type{
			Type:     apperror.UnAuthenticated,
			AppError: apperror.ErrInvalidCredential{},
		}
	}
	token, err := u.jwt.GenerateToken(fetchedUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}

func checkPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
