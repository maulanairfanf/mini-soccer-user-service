package error

import "errors"

const (
	Success = "success"
	Error   = "error"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("your requested item is not found")
	ErrorSQLError          = errors.New("something went wrong with the database")
	ErrorToManyRequest     = errors.New("too many request")
	ErrorBadRequest        = errors.New("bad request")
	ErrorUnauthorized      = errors.New("unauthorized")
	ErrorForbidden         = errors.New("forbidden")
	ErrorInvalidToken      = errors.New("invalid token")
	ErrorExpiredToken      = errors.New("expired token")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrNotFound,
	ErrorSQLError,
	ErrorToManyRequest,
	ErrorBadRequest,
	ErrorUnauthorized,
	ErrorForbidden,
	ErrorInvalidToken,
	ErrorExpiredToken,
}
