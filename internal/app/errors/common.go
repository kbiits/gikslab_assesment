package errors

import "net/http"

var ErrUnauthorized = &HttpError{
	StatusCode: http.StatusUnauthorized,
	Message:    "Unauthorized user",
}
