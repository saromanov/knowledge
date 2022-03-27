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

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var (
	errPageNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
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
