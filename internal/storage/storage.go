package storage

import (
	"context"

	model "github.com/saromanov/knowledge/internal/models/storage"
)
// Storage defines interface for storage
type Storage interface {
	Init(ctx context.Context) error
	CreatePage(ctx context.Context, m *model.Page) error
	CreateAuthor(ctx context.Context, m*model.Author) error
	GetPage(ctx context.Context, id int64) (*model.Page, error)
	Close(ctx context.Context) error
}

type StorageImpl[T Storage] struct {

}