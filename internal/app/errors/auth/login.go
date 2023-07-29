package errorAuth

import (
	"net/http"

	"github.com/kbiits/gikslab_assesment/internal/app/errors"
)

var ErrLoginInvalid = &errors.HttpError{
	Message:    "invalid login",
	StatusCode: http.StatusUnauthorized,
}
