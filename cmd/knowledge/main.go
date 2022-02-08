package main

import (
	"context"

	"github.com/oklog/run"
	"github.com/joeshaw/envdecode"
	"github.com/saromanov/knowledge/internal/storage/postgres"
	"github.com/saromanov/knowledge/internal/rest"
	"github.com/saromanov/knowledge/internal/service"
)

type config struct {
	Postgres postgres.Config
}
func main(){
	var cfg config
	if err := envdecode.StrictDecode(&cfg); err != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g := run.Group{}
	srv := service.New()
	pg := postgres.New(cfg.Postgres)
	if err := pg.Init(ctx); err != nil {
		panic(err)
	}
	defer pg.Close(ctx)
	r := rest.New(rest.Config{
		Address: "localhost:8044",
	}, pg)
	srv.Add(r, g)
	srv.Start(ctx)
}