package storage

import (
	"context"

	model "github.com/saromanov/knowledge/internal/models/storage"
)
// Storage defines interface for storage
type Storage interface {
	Init(ctx context.Context) error
	CreatePage(ctx context.Context, m *model.Page) (int64, error)
	CreateAuthor(ctx context.Context, m*model.Author) (int64, error)
	GetPage(ctx context.Context, id int64) (*model.Page, error)
	DeletePage(ctx context.Context, id int64) error
	GetPages(ctx context.Context, author string)([]*model.Page, error)
	Close(ctx context.Context) error
}

type StorageImpl[T Storage] struct {

}