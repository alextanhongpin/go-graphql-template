package external

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alextanhongpin/go-graphql-template/domain/account"
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/domain/user"
	"github.com/alextanhongpin/go-graphql-template/external/graph"
	"github.com/alextanhongpin/go-graphql-template/external/graph/schema"
	"github.com/alextanhongpin/go-graphql-template/external/middleware"
	"github.com/alextanhongpin/go-graphql-template/external/security"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func NewGraph(q entity.Querier, authz security.Authorizer, usersvc user.Service, accountsvc account.Service) http.Handler {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	r := graph.NewRoot()
	s := graphql.MustParseSchema(schema.String, r, opts...)

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))

	middlewares := []middleware.Middleware{
		middleware.CORS,
		middleware.ClientIP,
		middleware.Authz(authz),
		middleware.Querier(q), // TODO: Remove.
		middleware.DataLoader(usersvc),

		// Services.
		middleware.UserService(usersvc),
		middleware.AccountService(accountsvc),
	}
	mux.Handle("/query", middleware.Chain(&relay.Handler{Schema: s}, middlewares...))

	return mux
}
