package valueobject_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewCondition(t *testing.T) {
	t.Run("should return new condition", func(t *testing.T) {
		tests := map[string]struct {
			field     string
			operation valueobject.Operator
			value     string
			expected  *valueobject.Condition
		}{
			"valid": {
				field:     "a",
				operation: valueobject.Equal,
				value:     "b",
				expected:  &valueobject.Condition{Field: "a", Operation: "=", Value: "b"},
			},
			"ilike": {
				field:     "a",
				operation: valueobject.Ilike,
				value:     "b",
				expected:  &valueobject.Condition{Field: "a", Operation: "ILIKE", Value: "%b%"},
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				condition := valueobject.NewCondition(tt.field, tt.operation, tt.value)
				assert.Equal(t, tt.expected, condition)
			})
		}
	})
	t.Run("should return nil", func(t *testing.T) {
		tests := map[string]struct {
			field     string
			operation valueobject.Operator
			value     string
		}{
			"empty value": {field: "a", operation: "=", value: ""},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				condition := valueobject.NewCondition(tt.field, tt.operation, tt.value)
				assert.Nil(t, condition)
			})
		}
	})
}
