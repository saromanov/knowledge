package main

import (
	"context"

	"github.com/saromanov/knowledge/internal/rest"
)
func main(){
	ctx := context.Background()
	r := rest.New(rest.Config{
		Address: "localhost:8044",
	})
	r.Run(ctx)
}