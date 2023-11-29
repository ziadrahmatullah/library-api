package valueobject_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewQuery(t *testing.T) {
	expected := &valueobject.Query{
		Conditions: make([]*valueobject.Condition, 0),
	}

	query := valueobject.NewQuery()

	assert.Equal(t, expected, query)
}

func TestQuery_Condition(t *testing.T) {
	conditions := valueobject.NewCondition("id", valueobject.Equal, 1)
	expected := &valueobject.Query{
		Conditions: []*valueobject.Condition{conditions},
	}

	query := valueobject.NewQuery().Condition("id", valueobject.Equal, 1)

	assert.Equal(t, expected, query)
}

func TestQuery_Paginate(t *testing.T) {
	expected := &valueobject.Query{
		Page:       1,
		PerPage:    1,
		Conditions: make([]*valueobject.Condition, 0),
	}

	query := valueobject.NewQuery().Paginate(1, 1)

	assert.Equal(t, expected, query)
}

func TestQuery_Order(t *testing.T) {
	expected := &valueobject.Query{
		OrderedBy:  "A",
		Conditions: make([]*valueobject.Condition, 0),
	}

	query := valueobject.NewQuery().Order("A")

	assert.Equal(t, expected, query)
}

func TestQuery_Lock(t *testing.T) {
	expected := &valueobject.Query{
		Locked:     true,
		Conditions: make([]*valueobject.Condition, 0),
	}

	query := valueobject.NewQuery().Lock()

	assert.Equal(t, expected, query)
}
