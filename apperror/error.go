package apperror

import "fmt"

const (
	ErrUniqueValueConstraint = "23505"
)

type ErrAlreadyExist struct {
	Resource string
	Field    string
	Value    any
}

func (e ErrAlreadyExist) Error() string {
	return fmt.Sprintf("%s with %s '%s' already exist", e.Resource, e.Field, e.Value)
}

type ErrNotFound struct {
	Resource string
	Field    string
	Value    any
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s with %s: %v is not found", e.Resource, e.Field, e.Value)
}

type ErrEmptyStock struct {
	Resource string
}

func (e ErrEmptyStock) Error() string {
	return fmt.Sprintf("there is no stock for this %s", e.Resource)
}

type ErrBinding struct {
	ErrBinding error
}

func (e ErrBinding) Error() string {
	return e.ErrBinding.Error()
}
