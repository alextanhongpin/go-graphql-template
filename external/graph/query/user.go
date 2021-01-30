package query

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/domain/model"
	"github.com/alextanhongpin/go-graphql-template/external/graph/resolver"
	"github.com/alextanhongpin/go-graphql-template/external/session"
)

type UserQuery struct{}

// NewUserQuery returns a new query.
func NewUserQuery() *UserQuery {
	return &UserQuery{}
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

func (q *UserQuery) User(ctx context.Context, args UserArgs) (*resolver.UserResolver, error) {
	p, err := args.ToServiceFindUser()
	if err != nil {
		return nil, err
	}

	usersvc, err := session.UserService(ctx)
	if err != nil {
		return nil, err
	}

	user, err := usersvc.FindUser(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewUserResolver(user), nil
}
