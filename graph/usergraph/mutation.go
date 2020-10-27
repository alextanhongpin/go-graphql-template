package usergraph

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

func (m *Mutation) CreateUser(ctx context.Context, args CreateUserArgs) (*resolver.UserResolver, error) {
	if err := m.ctx.Validator.Struct(args.Input); err != nil {
		return nil, err
	}

	var (
		name  = args.Input.Name
		email = args.Input.Email
	)
	user, err := m.ctx.Repository.CreateUser(ctx, entity.CreateUserParams{
		Name:              name,
		Email:             model.NewNullString(email),
		PreferredUsername: name,
	})
	if err != nil {
		return nil, err
	}
	return &resolver.UserResolver{
		User: user,
		Ctx:  m.ctx,
	}, nil
}

func (m *Mutation) UpdateUser(ctx context.Context, args UpdateUserArgs) (*resolver.UserResolver, error) {
	if err := m.ctx.Validator.Struct(args.Input); err != nil {
		return nil, err
	}

	var (
		name = args.Input.Name
		id   = args.Input.ID
	)
	userID, err := uuid.Parse(string(id))
	if err != nil {
		return nil, err
	}

	user, err := m.ctx.Repository.UpdateUser(ctx, entity.UpdateUserParams{
		ID:   userID,
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return &resolver.UserResolver{
		User: user,
		Ctx:  m.ctx,
	}, nil
}

func (m *Mutation) DeleteUser(ctx context.Context, args DeleteUserArgs) (*resolver.UserResolver, error) {
	if err := m.ctx.Validator.Struct(args.Input); err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(string(args.Input.ID))
	if err != nil {
		return nil, err
	}

	user, err := m.ctx.Repository.DeleteUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &resolver.UserResolver{
		User: user,
		Ctx:  m.ctx,
	}, nil
}
