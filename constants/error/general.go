package error

const {
	Success = "success"
	Error = "error"
}

var {
	ErrInternalServerError = errors.New(text :"internal server error")
	ErrNotFound = errors.New(text: "your requested item is not found")
	ErrorSQLError = erros.New(text: "something went wrong with the database")
	ErrorToManyRequest = errors.New(text: "too many request")
	ErrorBadRequest = errors.New(text: "bad request")
	ErrorUnauthorized = errors.New(text: "unauthorized")
	ErrorForbidden = errors.New(text: "forbidden")
	ErrorInvalidToken = errors.New(text: "invalid token")
	ErrorExpiredToken = errors.New(text: "expired token")
}

var GeneralErrors = []error {
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