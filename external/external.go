package external

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/external/graph"
	"github.com/alextanhongpin/go-graphql-template/external/graph/schema"
	"github.com/alextanhongpin/go-graphql-template/external/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func NewGraph(q entity.Querier, authz middleware.Authorizer) http.Handler {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	r := graph.NewRoot()
	s := graphql.MustParseSchema(schema.String, r, opts...)

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))

	middlewares := []middleware.Middleware{
		middleware.CORS,
		middleware.ClientIP,
		middleware.Authz(authz),
		middleware.Querier(q),
		middleware.DataLoader(q),
	}
	mux.Handle("/query", middleware.Chain(&relay.Handler{Schema: s}, middlewares...))

	return mux
}
