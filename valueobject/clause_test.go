package valueobject_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewClause(t *testing.T) {
	t.Run("should return expected clause with no error", func(t *testing.T) {
		tests := map[string]struct {
			page    string
			perPage string
			order   string
			clause  *valueobject.Clause
		}{
			"positive":            {page: "2", perPage: "10", order: "", clause: &valueobject.Clause{Page: 2, PerPage: 10, Order: ""}},
			"turn page to 1 if 0": {page: "0", perPage: "10", order: "", clause: &valueobject.Clause{Page: 1, PerPage: 10, Order: ""}},
			"perPage 0":           {page: "0", perPage: "0", order: "", clause: &valueobject.Clause{Page: 1, PerPage: 0, Order: ""}},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				clause, err := valueobject.NewClause(tt.page, tt.perPage, tt.order)
				assert.Equal(t, tt.clause, clause)
				assert.NoError(t, err)
			})
		}
	})
	t.Run("should return error", func(t *testing.T) {
		tests := map[string]struct {
			page    string
			perPage string
			order   string
		}{
			"empty page":    {page: "", perPage: "10", order: ""},
			"empty perPage": {page: "0", perPage: "", order: ""},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				clause, err := valueobject.NewClause(tt.page, tt.perPage, tt.order)
				assert.Nil(t, clause)
				assert.Error(t, err)
			})
		}
	})
}
