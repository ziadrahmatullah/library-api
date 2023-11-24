package repository

import (
	"errors"
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Find(clause valueobject.Clause, conditions []valueobject.Condition) []*T
	First(conditions []valueobject.Condition) *T
	Create(t *T) (*T, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func (r *baseRepository[T]) Find(clause valueobject.Clause, conditions []valueobject.Condition) []*T {
	var ts []*T
	limit, offset, order := parseClause(clause)
	query := r.db.Model(ts)
	log.Println(conditions)
	for _, condition := range conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	query.
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&ts)
	return ts
}

func (r *baseRepository[T]) First(conditions []valueobject.Condition) *T {
	var t *T
	query := r.db.Model(t)
	for _, condition := range conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	err := query.First(&t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return t
}

func (r *baseRepository[T]) Create(t *T) (*T, error) {
	result := r.db.Create(t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}
