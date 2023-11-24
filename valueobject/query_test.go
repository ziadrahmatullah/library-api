package valueobject_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewQuery(t *testing.T) {
	t.Run("should return expected query with no error", func(t *testing.T) {
		tests := map[string]struct {
			page    string
			perPage string
			order   string
			query   *valueobject.Query
		}{
			"positive":            {page: "2", perPage: "10", order: "", query: &valueobject.Query{Page: 2, PerPage: 10, Order: ""}},
			"turn page to 1 if 0": {page: "0", perPage: "10", order: "", query: &valueobject.Query{Page: 1, PerPage: 10, Order: ""}},
			"perPage 0":           {page: "0", perPage: "0", order: "", query: &valueobject.Query{Page: 1, PerPage: 0, Order: ""}},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				query, err := valueobject.NewQuery(tt.page, tt.perPage, tt.order)
				assert.Equal(t, tt.query, query)
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
				query, err := valueobject.NewQuery(tt.page, tt.perPage, tt.order)
				assert.Nil(t, query)
				assert.Error(t, err)
			})
		}
	})
}
