package apperror

import "fmt"

const (
	ErrUniqueValueConstraint = "23505"
)

type ErrAlreadyExist struct {
	Resource string
	Field    string
	Value    string
}

func (e ErrAlreadyExist) Error() string {
	return fmt.Sprintf("%s with %s '%s' already exist", e.Resource, e.Field, e.Value)
}

type ErrNotFound struct {
}

func (e ErrNotFound) Error() string {
	return ""
}
