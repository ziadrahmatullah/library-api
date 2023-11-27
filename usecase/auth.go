package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

func (u *authUsecase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	emailCondition := *valueobject.NewCondition("email", valueobject.Equal, user.Email)
	emailQuery := valueobject.Query{
		Conditions: []valueobject.Condition{emailCondition},
	}
	fetchedUser := u.userRepo.First(ctx, emailQuery)
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

	phoneCondition := *valueobject.NewCondition("phone", valueobject.Equal, user.Phone)
	phoneQuery := valueobject.Query{
		Conditions: []valueobject.Condition{phoneCondition},
	}
	fetchedUser = u.userRepo.First(ctx, phoneQuery)
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

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}
