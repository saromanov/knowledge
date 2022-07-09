package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	restModel "github.com/saromanov/knowledge/internal/models/rest"
	storageModel "github.com/saromanov/knowledge/internal/models/storage"
	"github.com/saromanov/knowledge/internal/rest/response"
	"github.com/saromanov/knowledge/internal/storage"
	"github.com/sirupsen/logrus"
)

type DeletePageHandler struct {
	store storage.Storage
}

// NewDeletePageHandler provides init
func NewDeletePageHandler(st storage.Storage) *DeletePageHandler {
	return &DeletePageHandler{
		store: st,
	}
}

// Handle defines get request for the page
func (h GetPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var page *storageModel.Page
	pageValue := r.Context().Value("page")
	if pageValue != nil {
		page = pageValue.(*storageModel.Page)
	}
	out, err := json.Marshal(page)
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

// DeletePageCtx getting middleware page from db
func (h GetPageHandler) DeletePageCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			result *storageModel.Page
			err    error
		)

		ctx := r.Context()
		if id := chi.URLParam(r, "pageID"); id != "" {
			idParsed, err := strconv.ParseInt(id, 10, 32)
			if err != nil {
				render.Render(w, r, errPageNotFound)
				return
			}
			err = h.store.DeletePage(ctx, idParsed)
		} else {
			render.Render(w, r, errPageNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, errPageNotFound)
			return
		}

		ctxRes := context.WithValue(ctx, "page", result)
		next.ServeHTTP(w, r.WithContext(ctxRes))
	})
}
