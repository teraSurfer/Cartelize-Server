package service

import "context"

// DbService describes the service.
type DbService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	ConnectDB(ctx context.Context, s string) error
	GetDB(ctx context.Context, s string) error
	Init(ctx context.Context, s string) error
}

type basicDbService struct{}

func (b *basicDbService) ConnectDB(ctx context.Context, s string) (e0 error) {
	// TODO implement the business logic of ConnectDB
	return e0
}
func (b *basicDbService) GetDB(ctx context.Context, s string) (e0 error) {
	// TODO implement the business logic of GetDB
	return e0
}
func (b *basicDbService) Init(ctx context.Context, s string) (e0 error) {
	// TODO implement the business logic of Init
	return e0
}

// NewBasicDbService returns a naive, stateless implementation of DbService.
func NewBasicDbService() DbService {
	return &basicDbService{}
}

// New returns a DbService with all of the expected middleware wired in.
func New(middleware []Middleware) DbService {
	var svc DbService = NewBasicDbService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
