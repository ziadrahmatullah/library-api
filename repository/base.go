package repository

import (
	"errors"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Find(query valueobject.Query) []*T
	First(query valueobject.Query) *T
	Create(t *T) (*T, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func (r *baseRepository[T]) Find(q valueobject.Query) []*T {
	var ts []*T
	limit, offset := getPagination(q)
	query := r.db.Model(ts)
	for _, s := range q.With {
		query.Joins(s)
	}
	for _, condition := range q.Conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	query.
		Limit(limit).
		Offset(offset).
		Order(q.Order).
		Find(&ts)
	return ts
}

func (r *baseRepository[T]) First(q valueobject.Query) *T {
	conditions := q.Conditions
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
