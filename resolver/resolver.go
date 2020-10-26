package resolver

import (
	"database/sql"

	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/model"
	"github.com/alextanhongpin/graphql-server-starter/resolver/usergraph"
)

type UserQuery = usergraph.Query

type UserMutation = usergraph.Mutation

// Resolver represents the root for all queries and mutations.
type Resolver struct {
	*UserQuery
	*UserMutation
}

// Options represents the options for Resolver.
type Options struct {
	DB *sql.DB
}

// New returns a Resolver configured with the Options.
func New(opts Options) *Resolver {
	r := entity.New(opts.DB)
	ctx := &model.ResolverContext{
		Repository: r,
	}

	return &Resolver{
		UserQuery:    usergraph.NewQuery(ctx),
		UserMutation: usergraph.NewMutation(ctx),
	}
}
