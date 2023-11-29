package apperror

import (
	"net/http"
)

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "invalid credentials"
}

func NewInvalidCredentialsError() error {
	return NewClientError(&InvalidCredentialsError{}, http.StatusUnauthorized)
}
