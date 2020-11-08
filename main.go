package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/alextanhongpin/go-graphql-template/app"
	"github.com/alextanhongpin/go-graphql-template/entity"
	"github.com/alextanhongpin/go-graphql-template/graph"
	"github.com/alextanhongpin/go-graphql-template/pkg/middleware"
	"github.com/alextanhongpin/pkg/grace"
)

func main() {
	sg := grace.NewShutdownGroup()

	db := app.NewDB(sg)
	//cache := app.NewCache(app.NewRedisConfig(), sg)
	_ = app.NewLogger(sg)
	signer := app.NewSigner()
	app.NewTracer(sg)

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schemas, err := app.Loader("/schema")
	if err != nil {
		log.Fatal(err)
	}
	r := graph.NewResolver(graph.Options{
		DB: db,
	})
	schema := graphql.MustParseSchema(schemas, r, opts...)

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query",
		middleware.ClientIP(
			middleware.Cors(
				middleware.DataLoader(
					middleware.Authz(
						&relay.Handler{Schema: schema},
						signer,
					),
					entity.New(db),
				),
			),
		),
	)

	shutdown := grace.New(mux, 8080)
	sg.Add(shutdown)

	<-grace.Signal()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sg.Close(ctx)
}
