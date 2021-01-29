package external

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/external/graph"
	"github.com/alextanhongpin/go-graphql-template/external/graph/resolver"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func NewGraph(q entity.Querier, authz graph.Authorizer) http.Handler {
	schemas := graph.LoadSchema("github.com/alextanhongpin/go-graphql-template:/external/graph/schema")
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	r := resolver.New()
	schema := graphql.MustParseSchema(schemas, r, opts...)

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))

	middlewares := []graph.Middleware{
		graph.CORS,
		graph.ClientIPProvider,
		graph.BuildAuthz(authz),
		graph.BuildQuerierProvider(q),
		graph.BuildDataLoaderProvider(graph.NewDataLoaders(q)),
	}
	mux.Handle("/query", graph.Chain(&relay.Handler{Schema: schema}, middlewares...))

	return mux
}
