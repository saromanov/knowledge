package handlers

import (
	"encoding/json"
	"net/http"

	restModel "github.com/saromanov/knowledge/internal/models/rest"
	storageModel "github.com/saromanov/knowledge/internal/models/storage"
	"github.com/saromanov/knowledge/internal/storage"

	"github.com/go-chi/render"
)

type createPageHandler struct {
	store storage.Storage
}

func NewCreateArticleHandler(st storage.Storage) *createPageHandler {
	return &createPageHandler{
		store: st,
	}
}
func (h *createPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var st storageModel.Page
	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		return
	}
	if err := h.store.CreatePage(ctx, &st); err != nil {
		return
	}

	render.JSON(w, r, restModel.Response{})
}
