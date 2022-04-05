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
func (h GetPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	authorValue := r.FromValue("author")
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, restModel.Response{
		Data: (json.RawMessage)(out),
	})

}

// GetPageCtx getting middleware page from db
func (h GetPageHandler) GetPageCtx(next http.Handler) http.Handler {
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
			result, err = h.store.GetPage(ctx, idParsed)
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
