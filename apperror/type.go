package apperror

type HandlerErrType int

const (
	BadRequest HandlerErrType = iota + 1
	NotFound
	Conflict
	UnAuthenticated
	UnAuthorized
)

type Type struct {
	Type     HandlerErrType
	AppError error
}

func (t Type) Error() string {
	return t.AppError.Error()
}
