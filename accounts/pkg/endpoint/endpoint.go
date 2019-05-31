package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/terasurfer/Cartelize-Server/accounts/pkg/service"
)

// CreateCustomerRequest collects the request parameters for the CreateCustomer method.
type CreateCustomerRequest struct {
	Customer service.Customer `json:"customer"`
}

// CreateCustomerResponse collects the response parameters for the CreateCustomer method.
type CreateCustomerResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeCreateCustomerEndpoint returns an endpoint that invokes CreateCustomer on the service.
func MakeCreateCustomerEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequest)
		s0, e1 := s.CreateCustomer(ctx, req.Customer)
		return CreateCustomerResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateCustomerResponse) Failed() error {
	return r.E1
}

// CreateAdminRequest collects the request parameters for the CreateAdmin method.
type CreateAdminRequest struct {
	Admin service.Admin `json:"admin"`
}

// CreateAdminResponse collects the response parameters for the CreateAdmin method.
type CreateAdminResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeCreateAdminEndpoint returns an endpoint that invokes CreateAdmin on the service.
func MakeCreateAdminEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAdminRequest)
		s0, e1 := s.CreateAdmin(ctx, req.Admin)
		return CreateAdminResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateAdminResponse) Failed() error {
	return r.E1
}

// GetCustomerByIDRequest collects the request parameters for the GetCustomerByID method.
type GetCustomerByIDRequest struct {
	Id string `json:"id"`
}

// GetCustomerByIDResponse collects the response parameters for the GetCustomerByID method.
type GetCustomerByIDResponse struct {
	C0 service.Customer `json:"c0"`
	E1 error            `json:"e1"`
}

// MakeGetCustomerByIDEndpoint returns an endpoint that invokes GetCustomerByID on the service.
func MakeGetCustomerByIDEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByIDRequest)
		c0, e1 := s.GetCustomerByID(ctx, req.Id)
		return GetCustomerByIDResponse{
			C0: c0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetCustomerByIDResponse) Failed() error {
	return r.E1
}

// GetCustomerByEmailRequest collects the request parameters for the GetCustomerByEmail method.
type GetCustomerByEmailRequest struct {
	Email string `json:"email"`
}

// GetCustomerByEmailResponse collects the response parameters for the GetCustomerByEmail method.
type GetCustomerByEmailResponse struct {
	C0 service.Customer `json:"c0"`
	E1 error            `json:"e1"`
}

// MakeGetCustomerByEmailEndpoint returns an endpoint that invokes GetCustomerByEmail on the service.
func MakeGetCustomerByEmailEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByEmailRequest)
		c0, e1 := s.GetCustomerByEmail(ctx, req.Email)
		return GetCustomerByEmailResponse{
			C0: c0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetCustomerByEmailResponse) Failed() error {
	return r.E1
}

// GetAdminByIDRequest collects the request parameters for the GetAdminByID method.
type GetAdminByIDRequest struct {
	Id string `json:"id"`
}

// GetAdminByIDResponse collects the response parameters for the GetAdminByID method.
type GetAdminByIDResponse struct {
	A0 service.Admin `json:"a0"`
	E1 error         `json:"e1"`
}

// MakeGetAdminByIDEndpoint returns an endpoint that invokes GetAdminByID on the service.
func MakeGetAdminByIDEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAdminByIDRequest)
		a0, e1 := s.GetAdminByID(ctx, req.Id)
		return GetAdminByIDResponse{
			A0: a0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAdminByIDResponse) Failed() error {
	return r.E1
}

// GetAdminByEmailRequest collects the request parameters for the GetAdminByEmail method.
type GetAdminByEmailRequest struct {
	Email string `json:"email"`
}

// GetAdminByEmailResponse collects the response parameters for the GetAdminByEmail method.
type GetAdminByEmailResponse struct {
	A0 service.Admin `json:"a0"`
	E1 error         `json:"e1"`
}

// MakeGetAdminByEmailEndpoint returns an endpoint that invokes GetAdminByEmail on the service.
func MakeGetAdminByEmailEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAdminByEmailRequest)
		a0, e1 := s.GetAdminByEmail(ctx, req.Email)
		return GetAdminByEmailResponse{
			A0: a0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAdminByEmailResponse) Failed() error {
	return r.E1
}

// UpdateCustomerRequest collects the request parameters for the UpdateCustomer method.
type UpdateCustomerRequest struct {
	Customer service.Customer `json:"customer"`
}

// UpdateCustomerResponse collects the response parameters for the UpdateCustomer method.
type UpdateCustomerResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeUpdateCustomerEndpoint returns an endpoint that invokes UpdateCustomer on the service.
func MakeUpdateCustomerEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCustomerRequest)
		s0, e1 := s.UpdateCustomer(ctx, req.Customer)
		return UpdateCustomerResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateCustomerResponse) Failed() error {
	return r.E1
}

// UpdateAdminRequest collects the request parameters for the UpdateAdmin method.
type UpdateAdminRequest struct {
	Admin service.Admin `json:"admin"`
}

// UpdateAdminResponse collects the response parameters for the UpdateAdmin method.
type UpdateAdminResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeUpdateAdminEndpoint returns an endpoint that invokes UpdateAdmin on the service.
func MakeUpdateAdminEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAdminRequest)
		s0, e1 := s.UpdateAdmin(ctx, req.Admin)
		return UpdateAdminResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateAdminResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateCustomer implements Service. Primarily useful in a client.
func (e Endpoints) CreateCustomer(ctx context.Context, customer service.Customer) (s0 string, e1 error) {
	request := CreateCustomerRequest{Customer: customer}
	response, err := e.CreateCustomerEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateCustomerResponse).S0, response.(CreateCustomerResponse).E1
}

// CreateAdmin implements Service. Primarily useful in a client.
func (e Endpoints) CreateAdmin(ctx context.Context, admin service.Admin) (s0 string, e1 error) {
	request := CreateAdminRequest{Admin: admin}
	response, err := e.CreateAdminEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateAdminResponse).S0, response.(CreateAdminResponse).E1
}

// GetCustomerByID implements Service. Primarily useful in a client.
func (e Endpoints) GetCustomerByID(ctx context.Context, id string) (c0 service.Customer, e1 error) {
	request := GetCustomerByIDRequest{Id: id}
	response, err := e.GetCustomerByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetCustomerByIDResponse).C0, response.(GetCustomerByIDResponse).E1
}

// GetCustomerByEmail implements Service. Primarily useful in a client.
func (e Endpoints) GetCustomerByEmail(ctx context.Context, email string) (c0 service.Customer, e1 error) {
	request := GetCustomerByEmailRequest{Email: email}
	response, err := e.GetCustomerByEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetCustomerByEmailResponse).C0, response.(GetCustomerByEmailResponse).E1
}

// GetAdminByID implements Service. Primarily useful in a client.
func (e Endpoints) GetAdminByID(ctx context.Context, id string) (a0 service.Admin, e1 error) {
	request := GetAdminByIDRequest{Id: id}
	response, err := e.GetAdminByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAdminByIDResponse).A0, response.(GetAdminByIDResponse).E1
}

// GetAdminByEmail implements Service. Primarily useful in a client.
func (e Endpoints) GetAdminByEmail(ctx context.Context, email string) (a0 service.Admin, e1 error) {
	request := GetAdminByEmailRequest{Email: email}
	response, err := e.GetAdminByEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAdminByEmailResponse).A0, response.(GetAdminByEmailResponse).E1
}

// UpdateCustomer implements Service. Primarily useful in a client.
func (e Endpoints) UpdateCustomer(ctx context.Context, customer service.Customer) (s0 string, e1 error) {
	request := UpdateCustomerRequest{Customer: customer}
	response, err := e.UpdateCustomerEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateCustomerResponse).S0, response.(UpdateCustomerResponse).E1
}

// UpdateAdmin implements Service. Primarily useful in a client.
func (e Endpoints) UpdateAdmin(ctx context.Context, admin service.Admin) (s0 string, e1 error) {
	request := UpdateAdminRequest{Admin: admin}
	response, err := e.UpdateAdminEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateAdminResponse).S0, response.(UpdateAdminResponse).E1
}
