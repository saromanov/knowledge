package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
		if err := server.Shutdown(shutdownCtx); err != nil {
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
	getHandler := handlers.NewGetPageHandler(r.st)
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Route("/api/v1", func(ro chi.Router) {
		ro.Route("/pages", func(ros chi.Router) {
			ros.Post("/", handlers.NewCreateArticleHandler(r.st).Handle)
			ros.Route("/{pageID}", func(rop chi.Router) {
				rop.Get("/", getHandler.Handle)
			})
		})
		ro.Post("/authors", handlers.NewCreateAuthorHandler(r.st).Handle)
	})
	return router
}
