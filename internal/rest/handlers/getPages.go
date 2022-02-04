package handlers

import (
	"net/http"

	"github.com/saromanov/knowledge/internal/storage"
)

type GetPageHandler struct {
	store storage.Storage
}

// NewGetArticleHandler provides init
func NewGetArticleHandler(st storage.Storage) *GetPageHandler {
	return &GetPageHandler{
		store: st,
	}
}
func (h GetPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
