package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"gorm.io/gorm"
)

const (
	defaultLimit = 10
)

func getPagination(q *valueobject.Query) (int, int) {
	limit := q.PerPage
	if limit == 0 {
		limit = defaultLimit
	}
	offset := (q.Page - 1) * limit
	return limit, offset
}

type contextKeyTx string

func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, contextKeyTx("tx"), tx)
}

func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(contextKeyTx("tx")).(*gorm.DB); ok {
		return tx
	}
	return nil
}
