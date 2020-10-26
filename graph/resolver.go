package graph

import (
	"database/sql"

	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/graph/accountgraph"
	"github.com/alextanhongpin/graphql-server-starter/graph/usergraph"
	"github.com/alextanhongpin/graphql-server-starter/model"
)

type (
	UserQuery       = usergraph.Query
	UserMutation    = usergraph.Mutation
	AccountQuery    = accountgraph.Query
	AccountMutation = accountgraph.Mutation
)

// Resolver represents the root for all queries and mutations.
type Resolver struct {
	*AccountQuery
	*AccountMutation
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
		AccountQuery:    accountgraph.NewQuery(ctx),
		AccountMutation: accountgraph.NewMutation(ctx),
		UserQuery:       usergraph.NewQuery(ctx),
		UserMutation:    usergraph.NewMutation(ctx),
	}
}
