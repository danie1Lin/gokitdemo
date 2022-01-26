package handlers

import (
	"context"

	pb "trussdemo"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.LiveEventServer {
	return liveeventService{}
}

type liveeventService struct {
	repo Repository
}

func (s liveeventService) CreateEvent(ctx context.Context, in *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	var resp pb.CreateEventResponse
	event := protoToModel(in.Event)
	if event != nil {
		return nil, ErrInvalidEventArgument
	}
	if err := s.repo.CreateEvent(event); err != nil {
		return nil, err
	}
	resp.Event = modelToProto(event)
	return &resp, nil
}
