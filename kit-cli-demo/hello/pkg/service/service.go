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

type basicHelloService struct{}

func (b *basicHelloService) Foo(ctx context.Context, req FooRequest) (resp FooResponse, err error) {
	// TODO implement the business logic of Foo
	return resp, err
}

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService() HelloService {
	return &basicHelloService{}
}

// New returns a HelloService with all of the expected middleware wired in.
func New(middleware []Middleware) HelloService {
	var svc HelloService = NewBasicHelloService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
