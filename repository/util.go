package repository

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"

const (
	defaultLimit = 10
)

func getPagination(q valueobject.Query) (int, int) {
	limit := q.PerPage
	if limit == 0 {
		limit = defaultLimit
	}
	offset := (q.Page - 1) * limit
	return limit, offset
}
