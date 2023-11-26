package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	FindUserById(uint) (*models.User, error)
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

