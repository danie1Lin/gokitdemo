package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "hello/pkg/service"
)

// FooRequest collects the request parameters for the Foo method.
type FooRequest struct {
	Req service.FooRequest `json:"req"`
}

// FooResponse collects the response parameters for the Foo method.
type FooResponse struct {
	Resp service.FooResponse `json:"resp"`
	Err  error               `json:"err"`
}

// MakeFooEndpoint returns an endpoint that invokes Foo on the service.
func MakeFooEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		resp, err := s.Foo(ctx, req.Req)
		return FooResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r FooResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Foo implements Service. Primarily useful in a client.
func (e Endpoints) Foo(ctx context.Context, req service.FooRequest) (resp service.FooResponse, err error) {
	request := FooRequest{Req: req}
	response, err := e.FooEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FooResponse).Resp, response.(FooResponse).Err
}
