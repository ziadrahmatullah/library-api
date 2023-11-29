package apperror

import (
	"fmt"
	"net/http"
)

type ResourceNotFound struct {
	resource string
	field    string
	value    any
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s with %s: %v not found", e.resource, e.field, e.value)
}

func NewResourceNotFound(resource string, field string, value any) error {
	return NewClientError(&ResourceNotFound{
		resource: resource,
		field:    field,
		value:    value,
	}, http.StatusNotFound)
}
