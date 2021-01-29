package resolver

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/domain/model"
	"github.com/alextanhongpin/go-graphql-template/external/graph"
	"github.com/google/uuid"
)

type AccountQuery struct {
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
func NewAccountQuery() *AccountQuery {
	return &AccountQuery{}
}

func (q *AccountQuery) Account(ctx context.Context, args AccountArgs) (*AccountResolver, error) {
	accountID, err := uuid.Parse(string(args.ID))
	if err != nil {
		return nil, err
	}

	r, err := graph.Querier(ctx)
	if err != nil {
		return nil, err
	}

	account, err := r.FindAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return NewAccountResolver(account), nil
}
