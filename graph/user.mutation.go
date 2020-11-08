package graph

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/entity"
	"github.com/alextanhongpin/go-graphql-template/model"

	"github.com/google/uuid"
)

type UserMutation struct {
	ctx *Context
}

func NewUserMutation(ctx *Context) *UserMutation {
	return &UserMutation{ctx: ctx}
}

func (m *UserMutation) CreateUser(ctx context.Context, args CreateUserArgs) (*UserResolver, error) {
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
	return &UserResolver{
		user: user,
		ctx:  m.ctx,
	}, nil
}

func (m *UserMutation) UpdateUser(ctx context.Context, args UpdateUserArgs) (*UserResolver, error) {
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
	return &UserResolver{
		user: user,
		ctx:  m.ctx,
	}, nil
}

func (m *UserMutation) DeleteUser(ctx context.Context, args DeleteUserArgs) (*UserResolver, error) {
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
	return &UserResolver{
		user: user,
		ctx:  m.ctx,
	}, nil
}
