package graph

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/entity"

	"github.com/graph-gophers/graphql-go"
)

// UserResolver holds the user entity to resolve.
type UserResolver struct {
	user entity.User
	ctx  *Context
}

// NewUserResolver returns a new User resolver.
func NewUserResolver(ctx *Context, user entity.User) *UserResolver {
	return &UserResolver{
		user: user,
		ctx:  ctx,
	}
}

// ID returns the user's id.
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID.String())
}

// Name returns the user's name.
func (r *UserResolver) Name() string {
	return r.user.Name
}

// Email returns the user's unique email address.
func (r *UserResolver) Email() string {
	return r.user.Email.String
}

func (r *UserResolver) Accounts(ctx context.Context) ([]*AccountResolver, error) {
	userID := r.user.ID

	accounts, err := r.ctx.Repository.FindAccountsWithUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	result := make([]*AccountResolver, len(accounts))
	for idx, account := range accounts {
		result[idx] = &AccountResolver{
			account: account,
			ctx:     r.ctx,
		}
	}
	return result, nil
}
