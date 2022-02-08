package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	models "github.com/saromanov/knowledge/internal/models/storage"
	"github.com/saromanov/knowledge/internal/storage"
)

type postgres struct {
	cfg Config
	db *sql.DB
}

// New provides initialization of the module
func New(cfg Config) storage.Storage {
	return &postgres{
		cfg: cfg,
	}
}

func (p *postgres) CreatePage(ctx context.Context, m *models.Page) error {
	return nil
}

// connect provides connection to postgres
func (p *postgres) connect() error {
	var lastErr error
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.cfg.Host, p.cfg.Port, p.cfg.Username, p.cfg.Password, p.cfg.DB)
	for i := 0; i < p.cfg.ConnectRetries; i++ {
		db, err := sql.Open("postgres", conn)
		if err != nil {
			lastErr = err
			continue
		}
		if err := db.Ping(); err != nil {
			lastErr = err
		}
		p.db = db
		return nil
	}
	return lastErr
}
