package service

import (
	"context"
)

// HelloService describes the service.
type HelloService interface {
	// Add your methods here
	Foo(ctx context.Context, req FooRequest) (resp FooResponse, err error)
}

type FooRequest struct {
	name     string
	optional *string
}

type FooResponse struct {
}
