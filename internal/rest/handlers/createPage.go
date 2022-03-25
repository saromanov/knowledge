package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/saromanov/knowledge/internal/models/convert"
	restModel "github.com/saromanov/knowledge/internal/models/rest"
	"github.com/saromanov/knowledge/internal/rest/response"
	"github.com/saromanov/knowledge/internal/rest/validation"
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
		response.WriteError(w, r, http.StatusBadRequest, restModel.Error{
			Message: "unable to decode request",
		})
		return
	}
	if err := validation.PostPage(&st); err != nil {
		log.WithError(err).Error("unable to validate request")
		response.WriteError(w, r, http.StatusInternalServerError, restModel.Error{
			Message: "unable to validate request",
		})
		return
	}
	m := convert.RestPageToStoragePage(&st)
	m.CreatedAt = time.Now().UTC()
	id, err := h.store.CreatePage(ctx, m)
	if err != nil {
		log.WithError(err).Error("unable to create page")
		response.WriteError(w, r, http.StatusInternalServerError, restModel.Error{
			Message: "unable to create page",
		})
		return
	}

	st.ID = id
	out, err := json.Marshal(st)
	if err != nil {
		log.WithError(err).Error("unable to marshal response")
		response.WriteError(w, r, http.StatusInternalServerError, restModel.Error{
			Message: "unable to marshal response",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, restModel.Response{
		Data: (json.RawMessage)(out),
	})
}
