package service

import (
	"context"

	"github.com/oklog/run"
)

// Service defines interface for service
type Service interface {
	Run(ctx context.Context) error
	Close(ctx context.Context) error
}

// StartService provides starting of service
func StartService(ctx context.Context, srv Service, g *run.Group) error {
	g.Add(func() error {
		return srv.Run(ctx)
	}, func(err error) {
		srv.Close(ctx)
	})
	return nil
}
