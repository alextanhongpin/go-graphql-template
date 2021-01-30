package mutation

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/external/graph/resolver"
	"github.com/alextanhongpin/go-graphql-template/external/session"
)

type UserMutation struct {
}

func NewUserMutation() *UserMutation {
	return &UserMutation{}
}

func (m *UserMutation) CreateUser(ctx context.Context, args CreateUserArgs) (*resolver.UserResolver, error) {
	usersvc, err := session.UserService(ctx)
	if err != nil {
		return nil, err
	}

	u, err := usersvc.CreateUser(ctx, args.Input.ToServiceCreateUser())
	if err != nil {
		return nil, err
	}

	return resolver.NewUserResolver(u), nil
}

func (m *UserMutation) UpdateUser(ctx context.Context, args UpdateUserArgs) (*resolver.UserResolver, error) {
	p, err := args.Input.ToServiceUpdateUser()
	if err != nil {
		return nil, err
	}

	usersvc, err := session.UserService(ctx)
	if err != nil {
		return nil, err
	}

	u, err := usersvc.UpdateUser(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewUserResolver(u), nil
}

func (m *UserMutation) DeleteUser(ctx context.Context, args DeleteUserArgs) (*resolver.UserResolver, error) {
	p, err := args.Input.ToServiceDeleteUser()
	if err != nil {
		return nil, err
	}

	usersvc, err := session.UserService(ctx)
	if err != nil {
		return nil, err
	}

	u, err := usersvc.DeleteUser(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewUserResolver(u), nil
}
