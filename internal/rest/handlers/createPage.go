package handlers

import (
	"net/http"

	"github.com/saromanov/knowledge/internal/storage"
)

type CreatePageHandler struct {
	store storage.Storage
}

func NewCreateArticleHandler() *CreatePageHandler {
	return &CreatePageHandler{}
}
func (h CreatePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
