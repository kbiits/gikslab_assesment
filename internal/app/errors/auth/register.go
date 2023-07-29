package errorAuth

import (
	"net/http"

	errorType "github.com/kbiits/gikslab_assesment/internal/app/errors"
)

var ErrRegisterFailed = &errorType.HttpError{
	Message:    "Data cannot be processed",
	StatusCode: http.StatusUnprocessableEntity,
}

var ErrRegisterUnauthorized = &errorType.HttpError{
	Message:    "Unauthorized user",
	StatusCode: http.StatusUnauthorized,
}
