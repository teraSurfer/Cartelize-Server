package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/terasurfer/Cartelize-Server/db/pkg/service"
)

// ConnectDBRequest collects the request parameters for the ConnectDB method.
type ConnectDBRequest struct {
	S string `json:"s"`
}

// ConnectDBResponse collects the response parameters for the ConnectDB method.
type ConnectDBResponse struct {
	E0 error `json:"e0"`
}

// MakeConnectDBEndpoint returns an endpoint that invokes ConnectDB on the service.
func MakeConnectDBEndpoint(s service.DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConnectDBRequest)
		e0 := s.ConnectDB(ctx, req.S)
		return ConnectDBResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r ConnectDBResponse) Failed() error {
	return r.E0
}

// GetDBRequest collects the request parameters for the GetDB method.
type GetDBRequest struct {
	S string `json:"s"`
}

// GetDBResponse collects the response parameters for the GetDB method.
type GetDBResponse struct {
	E0 error `json:"e0"`
}

// MakeGetDBEndpoint returns an endpoint that invokes GetDB on the service.
func MakeGetDBEndpoint(s service.DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDBRequest)
		e0 := s.GetDB(ctx, req.S)
		return GetDBResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r GetDBResponse) Failed() error {
	return r.E0
}

// InitRequest collects the request parameters for the Init method.
type InitRequest struct {
	S string `json:"s"`
}

// InitResponse collects the response parameters for the Init method.
type InitResponse struct {
	E0 error `json:"e0"`
}

// MakeInitEndpoint returns an endpoint that invokes Init on the service.
func MakeInitEndpoint(s service.DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InitRequest)
		e0 := s.Init(ctx, req.S)
		return InitResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r InitResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// ConnectDB implements Service. Primarily useful in a client.
func (e Endpoints) ConnectDB(ctx context.Context, s string) (e0 error) {
	request := ConnectDBRequest{S: s}
	response, err := e.ConnectDBEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ConnectDBResponse).E0
}

// GetDB implements Service. Primarily useful in a client.
func (e Endpoints) GetDB(ctx context.Context, s string) (e0 error) {
	request := GetDBRequest{S: s}
	response, err := e.GetDBEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetDBResponse).E0
}

// Init implements Service. Primarily useful in a client.
func (e Endpoints) Init(ctx context.Context, s string) (e0 error) {
	request := InitRequest{S: s}
	response, err := e.InitEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(InitResponse).E0
}
