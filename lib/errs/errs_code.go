package errs

import "net/http"

var (
	GenericError = ErrContext{
		HttpStatus: http.StatusConflict,
		ErrMsg:     "Generic error, something went wrong",
	}

	BadRequest = ErrContext{
		HttpStatus: http.StatusBadRequest,
		ErrMsg:     "Bad request",
	}
)
