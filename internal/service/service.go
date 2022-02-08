package service

import (
	"context"
	"sync"

	"github.com/oklog/run"
)

// Service defines interface for service
type Service interface {
	Run(ctx context.Context) error
	Close(ctx context.Context) error
}

type ServiceImpl struct {
	mu   sync.RWMutex
	srvs []Service
}

// New creates serviceImpl
func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) Add(srv Service,g run.Group) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.srvs = append(s.srvs, srv)
	return nil
}

// Start provides starting of services
func (s *ServiceImpl) Start(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(s.srvs))
	for _, srv := range s.srvs {
		go func(srvImpl Service){
			defer wg.Done()
			if err := srvImpl.Run(ctx); err != nil {

			}
		}(srv)
	}
	wg.Wait()
	return nil
}
