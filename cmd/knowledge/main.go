package main

import (
	"context"

	"github.com/saromanov/knowledge/internal/storage"
	"github.com/saromanov/knowledge/internal/storage/postgres"
	"github.com/saromanov/knowledge/internal/rest"
)
func main(){
	ctx := context.Background()
	pg := postgres.New()
	r := rest.New(rest.Config{
		Address: "localhost:8044",
	})
	r.Run(ctx)
}