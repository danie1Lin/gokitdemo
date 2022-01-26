package handlers

import (
	"fmt"
	"trussdemo/errors"
)

var (
	ErrInvalidEventArgument           error = errors.WithStatusCode(fmt.Errorf("Invalid Event"), 400)
	ErrInvalidEventNameLengthArgument error = errors.WithStatusCode(fmt.Errorf("Invalid Event Name length"), 400)
)

func InvalidEventError(code int, reason string, location string) error {
	return errors.NewStandardError("Invalid Event", 400, code, reason, location)
}
