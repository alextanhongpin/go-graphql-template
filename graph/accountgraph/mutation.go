package accountgraph

import (
	"context"

	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/model"
	"github.com/alextanhongpin/graphql-server-starter/resolver"
	"github.com/google/uuid"
)

type Mutation struct {
	ctx *model.ResolverContext
}

func NewMutation(ctx *model.ResolverContext) *Mutation {
	return &Mutation{ctx: ctx}
}

func (m *Mutation) CreateAccount(ctx context.Context, args CreateAccountArgs) (*resolver.AccountResolver, error) {
	var (
		email    = args.Input.Email
		password = args.Input.Password
		userUUID = args.Input.UserID
	)
	userID, err := uuid.Parse(string(userUUID))
	if err != nil {
		return nil, err
	}

	account, err := m.ctx.Repository.CreateAccount(ctx, entity.CreateAccountParams{
		Uid:      email,
		Provider: entity.ProviderEmail,
		Token:    password,
		Email:    email,
		UserID:   userID,
	})
	if err != nil {
		return nil, err
	}
	return &resolver.AccountResolver{
		Account: account,
		Ctx:     m.ctx,
	}, nil
}

func (m *Mutation) UpdateAccount(ctx context.Context, args UpdateAccountArgs) (*resolver.AccountResolver, error) {
	var (
		email = args.Input.Email
		id    = args.Input.ID
	)
	accountID, err := uuid.Parse(string(id))
	if err != nil {
		return nil, err
	}
	account, err := m.ctx.Repository.UpdateAccount(ctx, entity.UpdateAccountParams{
		ID:    accountID,
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return &resolver.AccountResolver{
		Account: account,
		Ctx:     m.ctx,
	}, nil
}

func (m *Mutation) DeleteAccount(ctx context.Context, args DeleteAccountArgs) (*resolver.AccountResolver, error) {
	accountID, err := uuid.Parse(string(args.Input.ID))
	if err != nil {
		return nil, err
	}
	account, err := m.ctx.Repository.DeleteAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return &resolver.AccountResolver{
		Account: account,
		Ctx:     m.ctx,
	}, nil
}
