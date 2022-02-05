package postgres

import (
	"context"

	models "github.com/saromanov/knowledge/internal/models/storage"
	"github.com/saromanov/knowledge/internal/storage"
)

type postgres struct {
	cfg Config
}

// New provides initialization of the module
func New(cfg Config) storage.Storage {
	return &postgres {
		cfg: cfg,
	}
}

func (p *postgres) CreatePage(ctx context.Context, m *models.Page) error {
	return nil
}
// connect provides connection to postgres
func (p *postgres) connect() error {
	return nil
}