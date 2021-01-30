package mutation

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/external/graph/resolver"
	"github.com/alextanhongpin/go-graphql-template/external/session"
)

type AccountMutation struct {
}

func NewAccountMutation() *AccountMutation {
	return &AccountMutation{}
}

func (m *AccountMutation) CreateAccount(ctx context.Context, args CreateAccountArgs) (*resolver.AccountResolver, error) {
	p, err := args.Input.ToServiceCreateAccount()
	if err != nil {
		return nil, err
	}

	accountsvc, err := session.AccountService(ctx)
	if err != nil {
		return nil, err
	}

	account, err := accountsvc.CreateAccount(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewAccountResolver(account), nil
}

func (m *AccountMutation) UpdateAccount(ctx context.Context, args UpdateAccountArgs) (*resolver.AccountResolver, error) {
	p, err := args.Input.ToServiceUpdateAccount()
	if err != nil {
		return nil, err
	}

	accountsvc, err := session.AccountService(ctx)
	if err != nil {
		return nil, err
	}

	account, err := accountsvc.UpdateAccount(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewAccountResolver(account), nil
}

func (m *AccountMutation) DeleteAccount(ctx context.Context, args DeleteAccountArgs) (*resolver.AccountResolver, error) {
	p, err := args.Input.ToServiceDeleteAccount()
	if err != nil {
		return nil, err
	}

	accountsvc, err := session.AccountService(ctx)
	if err != nil {
		return nil, err
	}

	account, err := accountsvc.DeleteAccount(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewAccountResolver(account), nil
}
