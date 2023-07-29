package errorActivity

import (
	"net/http"

	"github.com/kbiits/gikslab_assesment/internal/app/errors"
)

var ErrFailedAddActivity = &errors.HttpError{
	Message: "Data cannot be processed",
	StatusCode: http.StatusUnprocessableEntity,
}

var ErrFailedUpdateActivity = &errors.HttpError{
	Message: "Data cannot be processed",
	StatusCode: http.StatusUnprocessableEntity,
}

var ErrFailedDeleteActivity = &errors.HttpError{
	Message: "Data cannot be processed",
	StatusCode: http.StatusUnprocessableEntity,
}