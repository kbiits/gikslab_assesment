package errorActivity

import (
	"net/http"

	"github.com/kbiits/gikslab_assesment/internal/app/errors"
)

var ErrListActivity = &errors.HttpError{
	Message:    "Data cannot be processed",
	StatusCode: http.StatusUnprocessableEntity,
}
