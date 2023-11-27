package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	FindUserByName(string) (users []models.User, err error)
	FindUserById(uint) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	NewUser(models.User)(*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindUsers() (users []models.User, err error) {
	err = u.db.Table("users").Find(&users).Error
	if err != nil {
		return nil, apperror.ErrFindUserQuery
	}
	return users, nil
}

func (b *userRepository) FindUserByName(name string) (users []models.User, err error) {
	err = b.db.Table("users").Where("name = ?", name).Find(&users).Error
	if err != nil {
		return nil, apperror.ErrFindUserQuery
	}
	return users, nil
}

func (u *userRepository) FindUserById(id uint) (user *models.User, err error) {
	result := u.db.Table("users").Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, apperror.ErrFindUserQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepository) FindByEmail(email string) (user *models.User, err error) {
	result := u.db.Table("users").Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, apperror.ErrFindUserQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepository) NewUser(user models.User) (newUser *models.User, err error){
	err = u.db.Table("users").Create(&user).Error
	if err != nil {
		return nil, apperror.ErrNewUserQuery
	}
	return &user, nil
}
