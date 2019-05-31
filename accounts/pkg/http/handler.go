package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "github.com/go-kit/kit/transport/http"
	endpoint "github.com/terasurfer/Cartelize-Server/accounts/pkg/endpoint"
	"net/http"
)

// makeCreateCustomerHandler creates the handler logic
func makeCreateCustomerHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-customer", http1.NewServer(endpoints.CreateCustomerEndpoint, decodeCreateCustomerRequest, encodeCreateCustomerResponse, options...))
}

// decodeCreateCustomerRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateCustomerRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateCustomerResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateCustomerResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateAdminHandler creates the handler logic
func makeCreateAdminHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-admin", http1.NewServer(endpoints.CreateAdminEndpoint, decodeCreateAdminRequest, encodeCreateAdminResponse, options...))
}

// decodeCreateAdminRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateAdminRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateAdminResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateAdminResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetCustomerByIDHandler creates the handler logic
func makeGetCustomerByIDHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-customer-by-id", http1.NewServer(endpoints.GetCustomerByIDEndpoint, decodeGetCustomerByIDRequest, encodeGetCustomerByIDResponse, options...))
}

// decodeGetCustomerByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetCustomerByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetCustomerByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetCustomerByIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetCustomerByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetCustomerByEmailHandler creates the handler logic
func makeGetCustomerByEmailHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-customer-by-email", http1.NewServer(endpoints.GetCustomerByEmailEndpoint, decodeGetCustomerByEmailRequest, encodeGetCustomerByEmailResponse, options...))
}

// decodeGetCustomerByEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetCustomerByEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetCustomerByEmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetCustomerByEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetCustomerByEmailResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAdminByIDHandler creates the handler logic
func makeGetAdminByIDHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-admin-by-id", http1.NewServer(endpoints.GetAdminByIDEndpoint, decodeGetAdminByIDRequest, encodeGetAdminByIDResponse, options...))
}

// decodeGetAdminByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAdminByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetAdminByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetAdminByIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAdminByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAdminByEmailHandler creates the handler logic
func makeGetAdminByEmailHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-admin-by-email", http1.NewServer(endpoints.GetAdminByEmailEndpoint, decodeGetAdminByEmailRequest, encodeGetAdminByEmailResponse, options...))
}

// decodeGetAdminByEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAdminByEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetAdminByEmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetAdminByEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAdminByEmailResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateCustomerHandler creates the handler logic
func makeUpdateCustomerHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-customer", http1.NewServer(endpoints.UpdateCustomerEndpoint, decodeUpdateCustomerRequest, encodeUpdateCustomerResponse, options...))
}

// decodeUpdateCustomerRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateCustomerRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateCustomerResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateCustomerResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateAdminHandler creates the handler logic
func makeUpdateAdminHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-admin", http1.NewServer(endpoints.UpdateAdminEndpoint, decodeUpdateAdminRequest, encodeUpdateAdminResponse, options...))
}

// decodeUpdateAdminRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateAdminRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateAdminResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateAdminResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
