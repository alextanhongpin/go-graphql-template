package graph

import (
	"database/sql"

	"github.com/alextanhongpin/go-graphql-template/entity"
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
func NewResolver(opts Options) *Resolver {
	ctx := &Context{
		Repository: entity.New(opts.DB),
	}

	return &Resolver{
		AccountQuery:    NewAccountQuery(ctx),
		AccountMutation: NewAccountMutation(ctx),
		UserQuery:       NewUserQuery(ctx),
		UserMutation:    NewUserMutation(ctx),
	}
}
