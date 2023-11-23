package repository

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"

const (
	defaultLimit = 10
)

func parseClause(c valueobject.Clause) (int, int, string) {
	limit := c.PerPage
	if limit == 0 {
		limit = defaultLimit
	}
	offset := (c.Page - 1) * limit
	return limit, offset, c.Order
}
