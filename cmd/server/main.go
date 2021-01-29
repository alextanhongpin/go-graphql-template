package main

import (
	"context"
	"time"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/external"
	"github.com/alextanhongpin/go-graphql-template/internal"
	"github.com/alextanhongpin/pkg/grace"
)

func main() {
	var sg grace.ShutdownGroup

	db, close := internal.NewDB()
	defer close()

	_, sync := internal.NewLogger(internal.NewLoggerConfig())
	defer sync()

	stop := internal.NewTracer()
	defer stop()

	q := entity.New(db)

	authz := internal.NewAuthorizer(internal.NewAuthorizerConfig())

	mux := external.NewGraph(q, authz)

	shutdown := grace.New(mux, 8080)
	sg.Add(shutdown)

	<-grace.Signal()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sg.Close(ctx)
}
