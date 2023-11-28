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
			query   *valueobject.Query
		}{
			"positive":            {page: "2", perPage: "10", query: &valueobject.Query{Page: 2, PerPage: 10, Conditions: make([]*valueobject.Condition, 0)}},
			"turn page to 1 if 0": {page: "0", perPage: "10", query: &valueobject.Query{Page: 1, PerPage: 10, Conditions: make([]*valueobject.Condition, 0)}},
			"perPage 0":           {page: "0", perPage: "0", query: &valueobject.Query{Page: 1, PerPage: 0, Conditions: make([]*valueobject.Condition, 0)}},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				query, err := valueobject.NewQuery(tt.page, tt.perPage)
				assert.Equal(t, tt.query, query)
				assert.NoError(t, err)
			})
		}
	})
	t.Run("should return error", func(t *testing.T) {
		tests := map[string]struct {
			page    string
			perPage string
		}{
			"empty page":    {page: "", perPage: "10"},
			"empty perPage": {page: "0", perPage: ""},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				query, err := valueobject.NewQuery(tt.page, tt.perPage)
				assert.Nil(t, query)
				assert.Error(t, err)
			})
		}
	})
}
