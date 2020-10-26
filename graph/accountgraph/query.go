package accountgraph

import (
	"context"

	"github.com/alextanhongpin/graphql-server-starter/model"
	"github.com/alextanhongpin/graphql-server-starter/resolver"
	"github.com/google/uuid"
)

type Query struct {
	ctx *model.ResolverContext
}

type AccountConnection struct {
	TotalCount int32
	Edges      []AccountEdge
	PageInfo   model.PageInfo
}

type AccountEdge struct {
	Cursor string
	Node   *resolver.AccountResolver
}

// NewQuery reutrns a new query.
func NewQuery(ctx *model.ResolverContext) *Query {
	return &Query{ctx}
}

func (q *Query) Account(ctx context.Context, args AccountArgs) (*resolver.AccountResolver, error) {
	accountID, err := uuid.Parse(string(args.ID))
	if err != nil {
		return nil, err
	}

	account, err := q.ctx.Repository.FindAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return &resolver.AccountResolver{
		Account: account,
		Ctx:     q.ctx,
	}, nil
}
