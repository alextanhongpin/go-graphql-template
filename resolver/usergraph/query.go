package usergraph

import (
	"context"
	"errors"

	"github.com/alextanhongpin/graphql-server-starter/model"
	"github.com/graph-gophers/graphql-go"
)

type Query struct {
	ctx *model.ResolverContext
}

type UserConnection struct {
	TotalCount int32
	Edges      []UserEdge
	//PageInfo   model.PageInfo
}

type UserEdge struct {
	Cursor string
	Node   *Resolver
}

func NewQuery(ctx *model.ResolverContext) *Query {
	return &Query{ctx}
}

type UserArgs struct {
	ID graphql.ID
}

func (q *Query) User(ctx context.Context, args UserArgs) (*Resolver, error) {
	return nil, errors.New("not implemented")
}
