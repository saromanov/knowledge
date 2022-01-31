package handlers

import (
	"net/http"

	"github.com/saromanov/knowledge/internal/storage"
)

type CreatePageHandler struct {
	store storage.Storage
}

func NewCreateArticleHandler(st storage.Storage) *CreatePageHandler {
	return &CreatePageHandler{
		store: st,
	}
}
func (h CreatePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
