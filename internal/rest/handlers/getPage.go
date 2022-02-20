package handlers

import (
	"net/http"

	"github.com/saromanov/knowledge/internal/storage"
)

type GetPageHandler struct {
	store storage.Storage
}

// NewGetPageHandler provides init
func NewGetPageHandler(st storage.Storage) *GetPageHandler {
	return &GetPageHandler{
		store: st,
	}
}
func (h GetPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	
}
