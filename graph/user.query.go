package graph

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/model"
	"github.com/google/uuid"
)

type UserQuery struct {
	ctx *Context
}

// NewUserQuery returns a new query.
func NewUserQuery(ctx *Context) *UserQuery {
	return &UserQuery{ctx}
}

type UserConnection struct {
	TotalCount int32
	Edges      []UserEdge
	PageInfo   model.PageInfo
}

type UserEdge struct {
	Cursor string
	Node   *UserResolver
}

func (q *UserQuery) User(ctx context.Context, args UserArgs) (*UserResolver, error) {
	userID, err := uuid.Parse(string(args.ID))
	if err != nil {
		return nil, err
	}

	user, err := q.ctx.Repository.FindUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &UserResolver{
		user: user,
		ctx:  q.ctx,
	}, nil
}
