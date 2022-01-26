package errors

type Error struct {
	error
	statusCode int
}

func (e *Error) StatusCode() int {
	return e.statusCode
}

func WithStatusCode(e error, code int) error {
	return &Error{error: e, statusCode: code}
}
