package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "hello/pkg/endpoint"
	pb "hello/pkg/grpc/pb"
)

// makeFooHandler creates the handler logic
func makeFooHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.FooEndpoint, decodeFooRequest, encodeFooResponse, options...)
}

// decodeFooResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Foo request.
// TODO implement the decoder
func decodeFooRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Hello' Decoder is not impelemented")
}

// encodeFooResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeFooResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Hello' Encoder is not impelemented")
}
func (g *grpcServer) Foo(ctx context.Context, req *pb.FooRequest) (*pb.FooReply, error) {
	_, rep, err := g.foo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FooReply), nil
}
