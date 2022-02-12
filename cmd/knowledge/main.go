package main

import (
	"context"

	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
	"github.com/joeshaw/envdecode"
	"github.com/saromanov/knowledge/internal/storage/postgres"
	"github.com/saromanov/knowledge/internal/rest"
	"github.com/saromanov/knowledge/internal/service"
)

type config struct {
	Postgres postgres.Config
}
func main(){
	logrus.SetFormatter(&logrus.JSONFormatter{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log := logrus.New().WithContext(ctx)
	var cfg config
	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.WithError(err).Fatal("unable to decode env variables")
		return
	}
	g := &run.Group{}
	pg := postgres.New(cfg.Postgres)
	log.Info("Trying to initialize postgres")
	if err := pg.Init(ctx); err != nil {
		log.WithError(err).Fatal("unable to init postgres")
	}
	log.Info("Postgres initialized")
	r := rest.New(rest.Config{
		Address: "localhost:8044",
	}, pg)

	if err := service.StartService(ctx, r, g); err != nil {
		log.WithError(err).Fatal("unable to start service")
	}
	log.Info("Finishing of the working")
}