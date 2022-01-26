package errors

import (
	"encoding/json"
	"fmt"
)

type Err struct {
	error
	statusCode int
}

func (e *Err) StatusCode() int {
	return e.statusCode
}

func WithStatusCode(e error, code int) *Err {
	return &Err{error: e, statusCode: code}
}

type StandardError struct {
	*Err   `json:"-"`
	Detail Detail `json:"detail,omitempty"`
}

// ref: https://google.github.io/styleguide/jsoncstyleguide.xml
type Detail struct {
	Code     int    `json:"code,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Location string `json:"location,omitempty"`
}

func NewStandardError(msg string, status int, code int, reason string, location string) error {
	return &StandardError{
		Err: WithStatusCode(fmt.Errorf(msg), status),
		Detail: Detail{
			Code:     code,
			Reason:   reason,
			Location: location,
		},
	}
}

func (e *StandardError) MarshalJSON() ([]byte, error) {
	data := struct {
		Message string `json:"message,omitempty"`
		Detail  Detail `json:"detail,omitempty"`
	}{
		Message: e.Error(),
		Detail:  e.Detail,
	}
	return json.Marshal(data)
}
