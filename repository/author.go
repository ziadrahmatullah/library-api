package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	BaseRepository[entity.Author]
}

type authorRepository struct {
	*baseRepository[entity.Author]
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{
		db:             db,
		baseRepository: &baseRepository[entity.Author]{db: db},
	}
}
