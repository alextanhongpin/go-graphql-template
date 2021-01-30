package query

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/domain/model"
	"github.com/alextanhongpin/go-graphql-template/external/graph/resolver"
	"github.com/alextanhongpin/go-graphql-template/external/session"
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
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
	Node   *resolver.AccountResolver
}

// AccountArgs represents the args for getting account.
type AccountArgs struct {
	ID graphql.ID
}

// NewAccountQuery reutrns a new query.
func NewAccountQuery() *AccountQuery {
	return &AccountQuery{}
}

func (q *AccountQuery) Account(ctx context.Context, args AccountArgs) (*resolver.AccountResolver, error) {
	accountID, err := uuid.Parse(string(args.ID))
	if err != nil {
		return nil, err
	}

	accountsvc, err := session.AccountService(ctx)
	if err != nil {
		return nil, err
	}

	account, err := accountsvc.FindAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return resolver.NewAccountResolver(account), nil
}
