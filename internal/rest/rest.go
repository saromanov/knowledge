package rest

import (
	"os"
	"os/signal"
	"context"
	"syscall"
	"net/http"

	"github.com/saromanov/knowledge/internal/rest/handlers"
	"github.com/saromanov/knowledge/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
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
	log := logrus.New().WithContext(ctx)
	server := &http.Server{Addr: r.cfg.Address, Handler: r.handlers()}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, _ := context.WithTimeout(serverCtx, r.cfg.ShutdownTimeout)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()
		if err := server.Shutdown(shutdownCtx);err != nil {
			log.WithError(err).Fatal("unable to shutdown service")
		}
		serverStopCtx()
	}()
	
	log.WithField("address", r.cfg.Address).Info("starting of the server")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.WithError(err).Error("closing of the server")
		return err
	}

	<-serverCtx.Done()
	return nil
}

func (r *rest) Close(ctx context.Context) error {
	return nil
}

func (r *rest) handlers() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Route("/api/v1", func(ro chi.Router) {
		ro.Post("/pages", handlers.NewCreateArticleHandler(r.st).Handle)
	})
	return router
}
