package repository

import (
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"gorm.io/gorm"
)

type BookRepository interface {
	Find(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book
	First(conditions []valueobject.Condition) *entity.Book
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

func (r *bookRepository) Find(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book {
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

func (r *bookRepository) First(conditions []valueobject.Condition) *entity.Book {
	var book *entity.Book
	query := r.db
	for _, condition := range conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	query.First(&book)
	return book
}

func (r *bookRepository) CreateBook(b *entity.Book) (*entity.Book, error) {
	result := r.db.Create(b)
	if result.Error != nil {
		return nil, result.Error
	}
	return b, nil
}
