package rest

import (
	"context"
	"net/http"

	"github.com/saromanov/knowledge/internal/rest/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type rest struct {
	cfg Config
}

func New(cfg Config) *rest {
	return &rest {
		cfg: cfg,
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
	router.Post("/pages", handlers.NewCreateArticleHandler())
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	if err := http.ListenAndServe(r.cfg.Address, router); err != nil {
		return err
	}
	return nil
} 