package apperror

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", ce.Code, ce.Message)
}

func (ce *CustomError) ToErrorRes() ErrorRes {
	return ErrorRes{
		Message: ce.Message,
	}
}

var (
	ErrFindBooksQuery       = NewCustomError(http.StatusInternalServerError, "find books query error")
	ErrFindBookByTitleQuery = NewCustomError(http.StatusInternalServerError, "find books by title query error")
	ErrFindBookByIdQuery    = NewCustomError(http.StatusInternalServerError, "find books by id query error")
	ErrNewBookQuery         = NewCustomError(http.StatusInternalServerError, "new book query error")

	ErrBookAlreadyExist = NewCustomError(http.StatusBadRequest, "book already exist")
	ErrBookNotFound     = NewCustomError(http.StatusBadRequest, "book not found")
	ErrBookOutOfStock   = NewCustomError(http.StatusBadRequest, "book out of stock")

	ErrFindUsersQuery    = NewCustomError(http.StatusInternalServerError, "find user query error")
	ErrFindUserByIdQuery = NewCustomError(http.StatusInternalServerError, "find user by id query error")
	ErrFindUserByName    = NewCustomError(http.StatusInternalServerError, "find user by name query error")
	ErrFindUserByEmail   = NewCustomError(http.StatusInternalServerError, "find user by email query error")
	ErrNewUserQuery      = NewCustomError(http.StatusInternalServerError, "new user query error")

	ErrUserNotFound     = NewCustomError(http.StatusBadRequest, "user not found")
	ErrEmailNotFound    = NewCustomError(http.StatusBadRequest, "email not found")
	ErrEmailALreadyUsed = NewCustomError(http.StatusBadRequest, "email already used")

	ErrUpdateQty            = NewCustomError(http.StatusInternalServerError, "update quantity error")
	ErrUpdateStatus         = NewCustomError(http.StatusInternalServerError, "update status error")
	ErrNewBorrowQuery       = NewCustomError(http.StatusInternalServerError, "new borrow query error")
	ErrFindBorrowQuery      = NewCustomError(http.StatusInternalServerError, "find borrow query error")
	ErrFindBorrowsQuery     = NewCustomError(http.StatusInternalServerError, "find borrows query error")
	ErrBorrowRecordNotFound = NewCustomError(http.StatusBadRequest, "borrow record not found")

	ErrGenerateJWTToken     = NewCustomError(http.StatusInternalServerError, "can't generate jwt token")
	ErrGenerateHashPassword = NewCustomError(http.StatusInternalServerError, "can't hash password")
	ErrMatchHashPassword    = NewCustomError(http.StatusBadRequest, "password doesn't match")
	ErrInvalidPassword      = NewCustomError(http.StatusBadRequest, "invalid password")

	ErrInvalidBody = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrTxCommit    = NewCustomError(http.StatusInternalServerError, "commit transaction error")
)
