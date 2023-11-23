package apperror

import "errors"

var (
	ErrBookAlreadyExist = errors.New("book already exist")
	ErrCannotBindJSON = errors.New("cannot bind json")
)