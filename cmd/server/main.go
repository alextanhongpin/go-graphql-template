package main

import (
	"context"
	"net/http"
	"time"

	"github.com/alextanhongpin/pkg/grace"
)

func main() {
	var sg grace.ShutdownGroup
	port := 8080

	r, stop := New()
	defer stop()
	sg.Add(run(r, port))

	<-grace.Signal()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sg.Close(ctx)
}

func run(r http.Handler, port int) func(context.Context) {
	return grace.New(r, port)
}
