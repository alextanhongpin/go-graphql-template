package usergraph

import (
	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/model"

	"github.com/graph-gophers/graphql-go"
)

type Resolver struct {
	user entity.User
	ctx  *model.ResolverContext
}

func NewResolver(ctx *model.ResolverContext, user entity.User) *Resolver {
	return &Resolver{
		user: user,
		ctx:  ctx,
	}
}

func (r *Resolver) ID() graphql.ID {
	return graphql.ID(r.user.ID.String())
}

func (r *Resolver) Name() string {
	return r.user.Name
}

func (r *Resolver) Email() string {
	return r.user.Email.String
}
