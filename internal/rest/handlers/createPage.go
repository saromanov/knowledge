package handlers

import (
	"encoding/json"
	"net/http"

	restModel "github.com/saromanov/knowledge/internal/models/rest"
	storageModel "github.com/saromanov/knowledge/internal/models/storage"
	"github.com/saromanov/knowledge/internal/storage"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
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
	log := logrus.New().WithContext(ctx)
	var st storageModel.Page
	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		log.WithError(err).Error("unable to parse request")
		return
	}
	if err := h.store.CreatePage(ctx, &st); err != nil {
		log.WithError(err).Error("unable to create page")
		return
	}

	render.JSON(w, r, restModel.Response{})
}
