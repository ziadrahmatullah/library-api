package apperror

import "net/http"

type MissingTokenError struct{}

func (e *MissingTokenError) Error() string {
	return "missing auth token"
}

func NewMissingTokenError() error {
	return NewClientError(&MissingTokenError{}, http.StatusUnauthorized)
}
