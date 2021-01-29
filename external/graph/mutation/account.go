package mutation

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/external/graph"
	"github.com/alextanhongpin/go-graphql-template/external/graph/resolver"
)

type AccountMutation struct {
}

func NewAccountMutation() *AccountMutation {
	return &AccountMutation{}
}

func (m *AccountMutation) CreateAccount(ctx context.Context, args CreateAccountArgs) (*resolver.AccountResolver, error) {
	p, err := args.Input.ToRepoCreateAccount()
	if err != nil {
		return nil, err
	}

	q, err := graph.Querier(ctx)
	if err != nil {
		return nil, err
	}

	account, err := q.CreateAccount(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewAccountResolver(account), nil
}

func (m *AccountMutation) UpdateAccount(ctx context.Context, args UpdateAccountArgs) (*resolver.AccountResolver, error) {
	p, err := args.Input.ToRepoUpdateAccount()
	if err != nil {
		return nil, err
	}

	q, err := graph.Querier(ctx)
	if err != nil {
		return nil, err
	}

	account, err := q.UpdateAccount(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewAccountResolver(account), nil
}

func (m *AccountMutation) DeleteAccount(ctx context.Context, args DeleteAccountArgs) (*resolver.AccountResolver, error) {
	p, err := args.Input.ToRepoDeleteAccount()
	if err != nil {
		return nil, err
	}

	q, err := graph.Querier(ctx)
	if err != nil {
		return nil, err
	}

	account, err := q.DeleteAccount(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewAccountResolver(account), nil
}
