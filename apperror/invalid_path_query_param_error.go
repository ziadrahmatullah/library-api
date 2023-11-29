package apperror

import "net/http"

type InvalidPathQueryParamError struct {
	err error
}

func (e *InvalidPathQueryParamError) Error() string {
	return e.err.Error()
}

func NewInvalidPathQueryParamError(err error) error {
	return NewClientError(&InvalidPathQueryParamError{err: err}, http.StatusBadRequest)
}
