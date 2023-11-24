package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type UserRepository interface{
	FindUsers() ([]models.User, error)
}

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindUsers() (users []models.User, err error) {
	table := u.db.Table("users")
	err = table.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
