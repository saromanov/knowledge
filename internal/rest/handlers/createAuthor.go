package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saromanov/knowledge/internal/models/convert"
	restModel "github.com/saromanov/knowledge/internal/models/rest"
	"github.com/saromanov/knowledge/internal/rest/response"
	"github.com/saromanov/knowledge/internal/storage"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type createAuthorHandler struct {
	store storage.Storage
}

// NewCreateAuthorHandler provides creating of the author
func NewCreateAuthorHandler(st storage.Storage) *createAuthorHandler {
	return &createAuthorHandler{
		store: st,
	}
}
func (h *createAuthorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logrus.New().WithContext(ctx)
	var st restModel.Author
	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		log.WithError(err).Error("unable to parse request")
		response.WriteError(w, r, http.StatusBadRequest, restModel.Error{
			Message: "unable to decode request",
		})
		return
	}
	m := convert.RestAuthorToStorageAuthor(&st)
	id, err := h.store.CreateAuthor(ctx, m)
	if err != nil {
		log.WithError(err).Error("unable to create author")
		response.WriteError(w, r, http.StatusInternalServerError, restModel.Error{
			Message: fmt.Sprintf("unable to create author: %v", err),
		})
		return
	}
	st.ID = fmt.Sprintf("%d", id)
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
