package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(DbService) DbService

type loggingMiddleware struct {
	logger log.Logger
	next   DbService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a DbService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next DbService) DbService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) ConnectDB(ctx context.Context, s string) (e0 error) {
	defer func() {
		l.logger.Log("method", "ConnectDB", "s", s, "e0", e0)
	}()
	return l.next.ConnectDB(ctx, s)
}
func (l loggingMiddleware) GetDB(ctx context.Context, s string) (e0 error) {
	defer func() {
		l.logger.Log("method", "GetDB", "s", s, "e0", e0)
	}()
	return l.next.GetDB(ctx, s)
}
func (l loggingMiddleware) Init(ctx context.Context, s string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Init", "s", s, "e0", e0)
	}()
	return l.next.Init(ctx, s)
}
