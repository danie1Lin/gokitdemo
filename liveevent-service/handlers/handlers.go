package handlers

import (
	"context"

	pb "trussdemo"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService(r Repository) pb.LiveEventServer {
	return liveeventService{
		repo: r,
	}
}

type liveeventService struct {
	repo Repository
}

func (s liveeventService) CreateEvent(ctx context.Context, in *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	var resp pb.CreateEventResponse
	if err := validateCreateEventRequest(in); err != nil {
		return nil, err
	}
	event := protoToModel(in.Event)
	if err := s.repo.CreateEvent(ctx, event); err != nil {
		return nil, err
	}
	resp.Event = modelToProto(event)
	return &resp, nil
}
