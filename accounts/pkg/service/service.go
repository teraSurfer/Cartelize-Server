package service

import (
	"context"
)

// AccountsService describes the service.
type AccountsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	CreateCustomer(ctx context.Context, customer Customer) (string, error)
	CreateAdmin(ctx context.Context, admin Admin) (string, error)
	GetCustomerByID(ctx context.Context, id string) (Customer, error)
	GetCustomerByEmail(ctx context.Context, email string) (Customer, error)
	GetAdminByID(ctx context.Context, id string) (Admin, error)
	GetAdminByEmail(ctx context.Context, email string) (Admin, error)
	UpdateCustomer(ctx context.Context, customer Customer) (string, error)
	UpdateAdmin(ctx context.Context, admin Admin) (string, error)
}

type basicAccountsService struct{}

func (b *basicAccountsService) CreateCustomer(ctx context.Context, customer Customer) (s0 string, e1 error) {
	// TODO implement the business logic of CreateCustomer
	return s0, e1
}
func (b *basicAccountsService) CreateAdmin(ctx context.Context, admin Admin) (s0 string, e1 error) {
	// TODO implement the business logic of CreateAdmin
	return s0, e1
}
func (b *basicAccountsService) GetCustomerByID(ctx context.Context, id string) (c0 Customer, e1 error) {
	// TODO implement the business logic of GetCustomerByID
	return c0, e1
}
func (b *basicAccountsService) GetCustomerByEmail(ctx context.Context, email string) (c0 Customer, e1 error) {
	// TODO implement the business logic of GetCustomerByEmail
	return c0, e1
}
func (b *basicAccountsService) GetAdminByID(ctx context.Context, id string) (a0 Admin, e1 error) {
	// TODO implement the business logic of GetAdminByID
	return a0, e1
}
func (b *basicAccountsService) GetAdminByEmail(ctx context.Context, email string) (a0 Admin, e1 error) {
	// TODO implement the business logic of GetAdminByEmail
	return a0, e1
}
func (b *basicAccountsService) UpdateCustomer(ctx context.Context, customer Customer) (s0 string, e1 error) {
	// TODO implement the business logic of UpdateCustomer
	return s0, e1
}
func (b *basicAccountsService) UpdateAdmin(ctx context.Context, admin Admin) (s0 string, e1 error) {
	// TODO implement the business logic of UpdateAdmin
	return s0, e1
}

// NewBasicAccountsService returns a naive, stateless implementation of AccountsService.
func NewBasicAccountsService() AccountsService {
	return &basicAccountsService{}
}

// New returns a AccountsService with all of the expected middleware wired in.
func New(middleware []Middleware) AccountsService {
	var svc AccountsService = NewBasicAccountsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
