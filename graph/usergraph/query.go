package usergraph

import (
	"context"

	"github.com/alextanhongpin/graphql-server-starter/model"
	"github.com/alextanhongpin/graphql-server-starter/resolver"
	"github.com/google/uuid"
)

type Query struct {
	ctx *model.ResolverContext
}

type UserConnection struct {
	TotalCount int32
	Edges      []UserEdge
	PageInfo   model.PageInfo
}

type UserEdge struct {
	Cursor string
	Node   *resolver.UserResolver
}

// NewQuery reutrns a new query.
func NewQuery(ctx *model.ResolverContext) *Query {
	return &Query{ctx}
}

func (q *Query) User(ctx context.Context, args UserArgs) (*resolver.UserResolver, error) {
	userID, err := uuid.Parse(string(args.ID))
	if err != nil {
		return nil, err
	}

	user, err := q.ctx.Repository.FindUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &resolver.UserResolver{
		User: user,
		Ctx:  q.ctx,
	}, nil
}
