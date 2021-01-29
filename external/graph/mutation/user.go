package mutation

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/external/graph"
	"github.com/alextanhongpin/go-graphql-template/external/graph/resolver"
)

type UserMutation struct {
}

func NewUserMutation() *UserMutation {
	return &UserMutation{}
}

func (m *UserMutation) CreateUser(ctx context.Context, args CreateUserArgs) (*resolver.UserResolver, error) {
	if err := graph.Validate(ctx, &args.Input); err != nil {
		return nil, err
	}

	q, err := graph.Querier(ctx)
	if err != nil {
		return nil, err
	}

	user, err := q.CreateUser(ctx, args.Input.ToRepoCreateUser())
	if err != nil {
		return nil, err
	}

	return resolver.NewUserResolver(user), nil
}

func (m *UserMutation) UpdateUser(ctx context.Context, args UpdateUserArgs) (*resolver.UserResolver, error) {
	if err := graph.Validate(ctx, &args.Input); err != nil {
		return nil, err
	}

	p, err := args.Input.ToRepoUpdateUser()
	if err != nil {
		return nil, err
	}

	q, err := graph.Querier(ctx)
	if err != nil {
		return nil, err
	}

	user, err := q.UpdateUser(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewUserResolver(user), nil
}

func (m *UserMutation) DeleteUser(ctx context.Context, args DeleteUserArgs) (*resolver.UserResolver, error) {
	if err := graph.Validate(ctx, &args.Input); err != nil {
		return nil, err
	}

	p, err := args.Input.ToRepoDeleteUser()
	if err != nil {
		return nil, err
	}

	q, err := graph.Querier(ctx)
	if err != nil {
		return nil, err
	}

	user, err := q.DeleteUser(ctx, p)
	if err != nil {
		return nil, err
	}

	return resolver.NewUserResolver(user), nil
}
