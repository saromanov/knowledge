package main

import (
	"context"

	"github.com/joeshaw/envdecode"
	"github.com/saromanov/knowledge/internal/storage/postgres"
	"github.com/saromanov/knowledge/internal/rest"
)

type config struct {
	Postgres postgres.Config
}
func main(){
	var cfg config
	if err := envdecode.StrictDecode(&cfg); err != nil {
		return
	}
	ctx := context.Background()
	pg := postgres.New(cfg.Postgres)
	r := rest.New(rest.Config{
		Address: "localhost:8044",
	}, pg)
	r.Run(ctx)
}