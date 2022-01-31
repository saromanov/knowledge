package postgres

import (
	"github.com/saromanov/knowledge/internal/storage"
)

type postgres struct {
	cfg Config
}

func new(cfg Config) storage.Storage {
	return &postgers {
		cfg: cfg,
	}
}