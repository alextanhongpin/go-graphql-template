package graph

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/model"
	"github.com/google/uuid"
)

type AccountQuery struct {
	ctx *Context
}

type AccountConnection struct {
	TotalCount int32
	Edges      []AccountEdge
	PageInfo   model.PageInfo
}

type AccountEdge struct {
	Cursor string
	Node   *AccountResolver
}

// NewAccountQuery reutrns a new query.
func NewAccountQuery(ctx *Context) *AccountQuery {
	return &AccountQuery{ctx}
}

func (q *AccountQuery) Account(ctx context.Context, args AccountArgs) (*AccountResolver, error) {
	accountID, err := uuid.Parse(string(args.ID))
	if err != nil {
		return nil, err
	}

	account, err := q.ctx.Repository.FindAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return &AccountResolver{
		account: account,
		ctx:     q.ctx,
	}, nil
}
