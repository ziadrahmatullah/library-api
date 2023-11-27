package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetAllUsers() ([]models.User, error)
	GetUserByName(string) ([]models.User, error)
	CreateUser(dto.RegisterReq) (*dto.RegisterRes, error)
	UserLogin(dto.LoginReq) (*dto.LoginRes, error)
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

func (u *userUsecase) GetUserByName(name string) ([]models.User, error) {
	return u.userRepository.FindUserByName(name)
}

func (u *userUsecase) CreateUser(registerData dto.RegisterReq) (data *dto.RegisterRes, err error) {
	user, _ := u.userRepository.FindByEmail(registerData.Email)
	if user != nil {
		return nil, apperror.ErrEmailALreadyUsed
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), 10)
	if err != nil {
		return nil, err
	}
	userModel := registerData.ToUserModelFromRegisterDTO(string(hashPassword))
	newUser, err := u.userRepository.NewUser(userModel)
	if err != nil {
		return nil, err
	}
	data = dto.ToUserResponsDTOFromModel(newUser)
	return data, nil
}

func (u *userUsecase) UserLogin(loginData dto.LoginReq) (token *dto.LoginRes, err error) {
	user, err := u.userRepository.FindByEmail(loginData.Email)
	if user == nil || err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return nil, err
	}
	newToken, err := dto.GenerateJWT(dto.JwtClaims{
		ID: user.ID,
	})
	if err != nil {
		return nil, err
	}
	return &dto.LoginRes{AccessToken: newToken}, nil
}
