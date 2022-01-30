package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	if err := http.ListenAndServe(r.cfg.Address, router); err != nil {
		return err
	}
	return nil
} 