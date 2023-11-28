package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers(context.Context) ([]models.User, error)
	FindUserByName(context.Context, string) (users []models.User, err error)
	FindUserById(context.Context, uint) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	NewUser(context.Context, models.User)(*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindUsers(ctx context.Context) (users []models.User, err error) {
	err = u.db.WithContext(ctx).Table("users").Find(&users).Error
	if err != nil {
		return nil, apperror.ErrFindUsersQuery
	}
	return users, nil
}

func (b *userRepository) FindUserByName(ctx context.Context, name string) (users []models.User, err error) {
	err = b.db.WithContext(ctx).Table("users").Where("name = ?", name).Find(&users).Error
	if err != nil {
		return nil, apperror.ErrFindUserByName
	}
	return users, nil
}

func (u *userRepository) FindUserById(ctx context.Context, id uint) (user *models.User, err error) {
	result := u.db.WithContext(ctx).Table("users").Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, apperror.ErrFindUserByIdQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (user *models.User, err error) {
	result := u.db.WithContext(ctx).Table("users").Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, apperror.ErrFindUserByEmail
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepository) NewUser(ctx context.Context, user models.User) (newUser *models.User, err error){
	err = u.db.WithContext(ctx).Table("users").Create(&user).Error
	if err != nil {
		return nil, apperror.ErrNewUserQuery
	}
	return &user, nil
}
