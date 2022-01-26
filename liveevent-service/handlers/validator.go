package handlers

import (
	pb "trussdemo"
)

func validateCreateEventRequest(e *pb.CreateEventRequest) error {
	if e == nil || e.Event == nil {
		return ErrInvalidEventArgument
	}
	if len(e.Event.Name) < 3 || len(e.Event.Name) > 20 {
		return ErrInvalidEventNameLengthArgument
	}
	return nil
}
