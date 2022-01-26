// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: fcd9ff140d
// Version Date: 2021-07-14T06:36:40Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "trussdemo"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC LiveEventServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.LiveEventServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// liveevent

		createevent: grpctransport.NewServer(
			endpoints.CreateEventEndpoint,
			DecodeGRPCCreateEventRequest,
			EncodeGRPCCreateEventResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the LiveEventServer interface
type grpcServer struct {
	createevent grpctransport.Handler
}

// Methods for grpcServer to implement LiveEventServer interface

func (s *grpcServer) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	_, rep, err := s.createevent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateEventResponse), nil
}

// Server Decode

// DecodeGRPCCreateEventRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createevent request to a user-domain createevent request. Primarily useful in a server.
func DecodeGRPCCreateEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateEventRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCCreateEventResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain createevent response to a gRPC createevent reply. Primarily useful in a server.
func EncodeGRPCCreateEventResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.CreateEventResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}