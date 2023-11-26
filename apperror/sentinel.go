package apperror

import "errors"

var (
	ErrBookAlreadyExist      = errors.New("book already exist")
	ErrCannotBindJSON        = errors.New("cannot bind json")
	ErrFindBooksQuery        = errors.New("find books query error")
	ErrFindBooksByTitleQuery = errors.New("find books by title query error")
	ErrNewBookQuery          = errors.New("new book query error")
	ErrBookNotFound          = errors.New("book not found")
	ErrBookOutOfStock        = errors.New("book out of stock")
	ErrFindUserQuery         = errors.New("find user query error")
	ErrFindUserByIdQuery     = errors.New("find user by id query error")
	ErrNewUserQuery          = errors.New("new user query error")
	ErrNewBorrowQuery        = errors.New("new borrow query error")
	ErrUpdateBookQtyQuery    = errors.New("update book qty query error")
	ErrUserNotFound          = errors.New("user not found")
	ErrTxCommit              = errors.New("commit transaction error")
	ErrFindBorrowQuery       = errors.New("find borrow query error")
	ErrBorrowRecordNotFound   = errors.New("borrow record not found")
)
