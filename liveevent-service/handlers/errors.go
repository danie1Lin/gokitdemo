package handlers

import (
	"fmt"
	"trussdemo/errors"
)

var (
	ErrInvalidEventArgument           = errors.WithStatusCode(fmt.Errorf("Invalid Event"), 400)
	ErrInvalidEventNameLengthArgument = errors.WithStatusCode(fmt.Errorf("Invalid Event Name length"), 400)
)
