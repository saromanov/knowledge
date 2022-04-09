package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	restModel "github.com/saromanov/knowledge/internal/models/rest"
	"github.com/saromanov/knowledge/internal/storage"
	"github.com/saromanov/knowledge/internal/rest/response"
	"github.com/sirupsen/logrus"
)


type GetPagesHandler struct {
	store storage.Storage
}

// NewGetPagesHandler provides init
func NewGetPagesHandler(st storage.Storage) *GetPagesHandler {
	return &GetPagesHandler{
		store: st,
	}
}

// Handle defines get request for get pages
func (h GetPagesHandler) Handle(w http.ResponseWriter, r *http.Request) {
	authorValue := r.URL.Query().Get("author")
	pages, err := h.store.GetPages(r.Context(), authorValue)
	if err != nil {
		logrus.WithError(err).Error("unable to get pages")
		response.WriteError(w, r, http.StatusInternalServerError, restModel.Error{
			Message: "unable to get pages",
		})
		render.Render(w, r, errPageNotFound)
		return
	}
	out, err := json.Marshal(pages)
	if err != nil {
		logrus.WithError(err).Error("unable to marshal response")
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
