package repository

import (
	"errors"
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book
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

func (r *bookRepository) FindAll(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book {
	var books []*entity.Book
	limit, offset, order := parseClause(clause)
	query := r.db.Joins("Author")
	log.Println(conditions)
	for _, condition := range conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	query.
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&books)
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
