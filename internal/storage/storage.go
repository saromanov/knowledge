package storage

import (
	"context"

	model "github.com/saromanov/knowledge/internal/models/storage"
)
// Storage defines interface for storage
type Storage interface {
	CreatePage(ctx context.Context, *model.Page) error
}