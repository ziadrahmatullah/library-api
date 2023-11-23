package repository

import (
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(c valueobject.Clause) []*entity.Book
	FindByTitle(name string) []*entity.Book
	CreateBook(book *entity.Book) (*entity.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) FindAll(c valueobject.Clause) []*entity.Book {
	var books []*entity.Book
	limit, offset, order := parseClause(c)
	r.db.
		Joins("Author").
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&books)
	return books
}

func (r *bookRepository) FindByTitle(title string) []*entity.Book {
	var books []*entity.Book
	r.db.Joins("Author").Where("title ILIKE ?", "%"+title+"%").Find(&books)
	return books
}

func (r *bookRepository) CreateBook(b *entity.Book) (*entity.Book, error) {
	result := r.db.Create(b)
	if result.Error != nil {
		var err *pgconn.PgError
		if errors.As(result.Error, &err) {
			if err.Code == apperror.ErrUniqueValueConstraint {
				return nil, apperror.ErrAlreadyExist{
					Resource: "book",
					Field:    "title",
					Value:    b.Title,
				}
			}
		}
		return nil, result.Error
	}
	return b, nil
}
