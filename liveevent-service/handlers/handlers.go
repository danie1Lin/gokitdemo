package handlers

import (
	"context"

	pb "trussdemo"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.LiveEventServer {
	return liveeventService{}
}

type liveeventService struct{}

func (s liveeventService) CreateEvent(ctx context.Context, in *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	var resp pb.CreateEventResponse
	return &resp, nil
}
