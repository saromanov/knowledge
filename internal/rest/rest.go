package rest

import (
	"context"
	"net/http"

	"github.com/saromanov/knowledge/internal/rest/handlers"
	"github.com/saromanov/knowledge/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type rest struct {
	cfg Config
	st  storage.Storage
}

func New(cfg Config, st storage.Storage) *rest {
	return &rest{
		cfg: cfg,
		st:  st,
	}
}

// Run starts of the server
func (r *rest) Run(ctx context.Context) error {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Route("api/v1", func(ro chi.Router) {
		ro.Post("/pages", handlers.NewCreateArticleHandler(r.st).Handle)
	})
	if err := http.ListenAndServe(r.cfg.Address, router); err != nil {
		return err
	}
	return nil
}
