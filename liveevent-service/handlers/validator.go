package handlers

import (
	pb "github.com/danie1Lin/gokitdemo"
)

func validateCreateEventRequest(e *pb.CreateEventRequest) error {
	if e == nil || e.Event == nil {
		return ErrInvalidEventArgument
	}
	if len(e.Event.Name) < 3 || len(e.Event.Name) > 20 {
		return InvalidEventError(1, "length must between 3 and 20", "name")
	}
	return nil
}
