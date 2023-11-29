package apperror

import "net/http"

type InvalidTokenError struct{}

func (e *InvalidTokenError) Error() string {
	return "invalid token"
}

func NewInvalidTokenError() error {
	return NewClientError(&InvalidTokenError{}, http.StatusUnauthorized)
}
