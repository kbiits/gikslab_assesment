package errors

type HttpError struct {
	StatusCode int
	Message    string
}

var _ error = (*HttpError)(nil)

func (e *HttpError) Error() string {
	return e.Message
}
