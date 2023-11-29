package apperror

import (
	"fmt"
	"net/http"
)

type UnavailableResourceError struct {
	resource string
}

func (e *UnavailableResourceError) Error() string {
	return fmt.Sprintf("there is no %s available", e.resource)
}

func NewUnavailableResourceError(resource string) error {
	return NewClientError(&UnavailableResourceError{
		resource: resource,
	}, http.StatusBadRequest)
}
