package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[entity.User]
}

type userRepository struct {
	*baseRepository[entity.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db:             db,
		baseRepository: &baseRepository[entity.User]{db: db},
	}
}
