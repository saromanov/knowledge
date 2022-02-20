package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/saromanov/knowledge/internal/models/convert"
	restModel "github.com/saromanov/knowledge/internal/models/rest"
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
	var st restModel.Page
	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		log.WithError(err).Error("unable to parse request")
		return
	}
	m := convert.RestPageToStoragePage(&st)
	if err := h.store.CreatePage(ctx, m); err != nil {
		log.WithError(err).Error("unable to create page")
		return
	}

	out, err := json.Marshal(st)
	if err != nil {
		log.WithError(err).Error("unable to marshal response")
		return
	}
	render.JSON(w, r, restModel.Response{
		Data: (json.RawMessage)(out),
	})
}
