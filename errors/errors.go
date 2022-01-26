package errors

type Error struct {
	error
	statusCode int
}

func WithStatusCode(e error, code int) error {
	return &Error{error: e, statusCode: code}
}
