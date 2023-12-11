package apperror

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return ce.Message
}

func (ce *CustomError) ToErrorRes() ErrorRes {
	return ErrorRes{
		Message: ce.Message,
	}
}

func (ce *CustomError) ToGrpcError() error {
	errorMapper := map[int]codes.Code{
		http.StatusInternalServerError: codes.Internal,
		http.StatusBadRequest:          codes.InvalidArgument,
		http.StatusUnauthorized:        codes.PermissionDenied,
	}
	return status.Error(errorMapper[ce.Code], ce.Message)
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

	ErrGenerateJWTToken       = NewCustomError(http.StatusInternalServerError, "can't generate jwt token")
	ErrGenerateHashPassword   = NewCustomError(http.StatusInternalServerError, "can't hash password")
	ErrInvalidPasswordOrEmail = NewCustomError(http.StatusBadRequest, "invalid password or email")
	ErrNotAuthorize           = NewCustomError(http.StatusUnauthorized, "not authorize")
	ErrInvalidAuthHeader      = NewCustomError(http.StatusUnauthorized, "erorr invalid authorization header")
	ErrInvalidJWTToken        = NewCustomError(http.StatusUnauthorized, "invalid jwt token")

	ErrInvalidBody = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrTxCommit    = NewCustomError(http.StatusInternalServerError, "commit transaction error")
)
