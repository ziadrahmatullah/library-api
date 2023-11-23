package apperror

import "errors"

var (
	ErrBookAlreadyExist = errors.New("book already exist")
	ErrCannotBindJSON = errors.New("cannot bind json")
	ErrFindBooksQuery = errors.New("find books query error")
	ErrFindBooksByTitleQuery = errors.New("find books by title query error")
	ErrNewBookQuery = errors.New("new book query error")
)