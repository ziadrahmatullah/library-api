package repository

import (
	"context"
	"errors"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository[T any] interface {
	UnitOfWork
	Find(ctx context.Context, query *valueobject.Query) ([]*T, error)
	First(ctx context.Context, query *valueobject.Query) (*T, error)
	Create(ctx context.Context, t *T) (*T, error)
	Update(ctx context.Context, t *T) (*T, error)
	Delete(ctx context.Context, t *T) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func (r *baseRepository[T]) Run(ctx context.Context, runner func(c context.Context) error) error {
	tx := r.db.Begin()
	ctx = injectTx(ctx, tx)
	err := runner(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *baseRepository[T]) conn(ctx context.Context) *gorm.DB {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.WithContext(ctx)
	}
	return r.db.WithContext(ctx)
}

func (r *baseRepository[T]) Find(ctx context.Context, q *valueobject.Query) ([]*T, error) {
	var ts []*T
	limit, offset := getPagination(q)
	query := r.conn(ctx).Model(ts)
	for _, s := range q.With {
		query.Joins(s)
	}
	for _, condition := range q.Conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	err := query.
		Limit(limit).
		Offset(offset).
		Order(q.OrderedBy).
		Find(&ts).
		Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (r *baseRepository[T]) First(ctx context.Context, q *valueobject.Query) (*T, error) {
	conditions := q.Conditions
	var t *T
	query := r.conn(ctx).Model(t)
	if q.Locked {
		query.Clauses(clause.Locking{Strength: "UPDATE"})
	}
	for _, condition := range conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	err := query.First(&t).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return t, nil
}

func (r *baseRepository[T]) Create(ctx context.Context, t *T) (*T, error) {
	result := r.conn(ctx).Create(t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

func (r *baseRepository[T]) Update(ctx context.Context, t *T) (*T, error) {
	result := r.conn(ctx).Model(t).Clauses(clause.Returning{}).Select("*").Updates(t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}
func (r *baseRepository[T]) Delete(ctx context.Context, t *T) error {
	result := r.conn(ctx).Delete(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
