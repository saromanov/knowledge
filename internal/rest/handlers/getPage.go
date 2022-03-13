package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	restModel "github.com/saromanov/knowledge/internal/models/rest"
	"github.com/saromanov/knowledge/internal/rest/response"
	"github.com/saromanov/knowledge/internal/storage"
	"github.com/sirupsen/logrus"
	"github.com/go-chi/render"
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

// Handle defines get request for the page
func (h GetPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logrus.New().WithContext(ctx)
	id := chi.URLParam(r, "id")
	if id == "" {
		log.Error("id param is not defined")
		response.WriteError(w, r, http.StatusBadRequest, restModel.Error{
			Message: "id param is not defined",
		})
	}

	idParsed, err := strconv.ParseInt(id, 10,32)
	if err != nil {
		log.WithError(err).Error("unable to parse id")
		response.WriteError(w, r, http.StatusBadRequest, restModel.Error{
			Message: "unable to parse id",
		})
	}
	result, err := h.store.GetPage(ctx, idParsed)
	if err != nil {
		log.WithError(err).Error("unable to get page by id")
		response.WriteError(w, r, http.StatusBadRequest, restModel.Error{
			Message: "unable to get page",
		})
	}

	out, err := json.Marshal(result)
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
