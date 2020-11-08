package graph

import (
	"context"

	"github.com/google/uuid"

	"github.com/alextanhongpin/go-graphql-template/entity"
)

type AccountMutation struct {
	ctx *Context
}

func NewAccountMutation(ctx *Context) *AccountMutation {
	return &AccountMutation{ctx: ctx}
}

func (m *AccountMutation) CreateAccount(ctx context.Context, args CreateAccountArgs) (*AccountResolver, error) {
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
	return &AccountResolver{
		account: account,
		ctx:     m.ctx,
	}, nil
}

func (m *AccountMutation) UpdateAccount(ctx context.Context, args UpdateAccountArgs) (*AccountResolver, error) {
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
	return &AccountResolver{
		account: account,
		ctx:     m.ctx,
	}, nil
}

func (m *AccountMutation) DeleteAccount(ctx context.Context, args DeleteAccountArgs) (*AccountResolver, error) {
	accountID, err := uuid.Parse(string(args.Input.ID))
	if err != nil {
		return nil, err
	}
	account, err := m.ctx.Repository.DeleteAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return &AccountResolver{
		account: account,
		ctx:     m.ctx,
	}, nil
}
