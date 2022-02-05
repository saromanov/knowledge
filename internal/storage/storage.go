package storage

import (
	"context"

	model "github.com/saromanov/knowledge/internal/models/storage"
)
// Storage defines interface for storage
type Storage interface {
	CreatePage(ctx context.Context, m *model.Page) error
}

type StorageImpl[T Storage] struct {

}