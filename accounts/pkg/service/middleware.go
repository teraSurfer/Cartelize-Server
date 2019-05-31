package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AccountsService) AccountsService

type loggingMiddleware struct {
	logger log.Logger
	next   AccountsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AccountsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AccountsService) AccountsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateCustomer(ctx context.Context, customer Customer) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "CreateCustomer", "customer", customer, "s0", s0, "e1", e1)
	}()
	return l.next.CreateCustomer(ctx, customer)
}
func (l loggingMiddleware) CreateAdmin(ctx context.Context, admin Admin) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "CreateAdmin", "admin", admin, "s0", s0, "e1", e1)
	}()
	return l.next.CreateAdmin(ctx, admin)
}
func (l loggingMiddleware) GetCustomerByID(ctx context.Context, id string) (c0 Customer, e1 error) {
	defer func() {
		l.logger.Log("method", "GetCustomerByID", "id", id, "c0", c0, "e1", e1)
	}()
	return l.next.GetCustomerByID(ctx, id)
}
func (l loggingMiddleware) GetCustomerByEmail(ctx context.Context, email string) (c0 Customer, e1 error) {
	defer func() {
		l.logger.Log("method", "GetCustomerByEmail", "email", email, "c0", c0, "e1", e1)
	}()
	return l.next.GetCustomerByEmail(ctx, email)
}
func (l loggingMiddleware) GetAdminByID(ctx context.Context, id string) (a0 Admin, e1 error) {
	defer func() {
		l.logger.Log("method", "GetAdminByID", "id", id, "a0", a0, "e1", e1)
	}()
	return l.next.GetAdminByID(ctx, id)
}
func (l loggingMiddleware) GetAdminByEmail(ctx context.Context, email string) (a0 Admin, e1 error) {
	defer func() {
		l.logger.Log("method", "GetAdminByEmail", "email", email, "a0", a0, "e1", e1)
	}()
	return l.next.GetAdminByEmail(ctx, email)
}
func (l loggingMiddleware) UpdateCustomer(ctx context.Context, customer Customer) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "UpdateCustomer", "customer", customer, "s0", s0, "e1", e1)
	}()
	return l.next.UpdateCustomer(ctx, customer)
}
func (l loggingMiddleware) UpdateAdmin(ctx context.Context, admin Admin) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "UpdateAdmin", "admin", admin, "s0", s0, "e1", e1)
	}()
	return l.next.UpdateAdmin(ctx, admin)
}
