package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	BaseRepository[entity.Book]
}

type bookRepository struct {
	*baseRepository[entity.Book]
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db:             db,
		baseRepository: &baseRepository[entity.Book]{db: db},
	}
}
