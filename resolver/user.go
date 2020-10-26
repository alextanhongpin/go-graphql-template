package resolver

import (
	"context"

	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/model"

	"github.com/graph-gophers/graphql-go"
)

// UserResolver holds the user entity to resolve.
type UserResolver struct {
	User entity.User
	Ctx  *model.ResolverContext
}

// NewUserResolver returns a new User resolver.
func NewUserResolver(ctx *model.ResolverContext, user entity.User) *UserResolver {
	return &UserResolver{
		User: user,
		Ctx:  ctx,
	}
}

// ID returns the user's id.
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.User.ID.String())
}

// Name returns the user's name.
func (r *UserResolver) Name() string {
	return r.User.Name
}

// Email returns the user's unique email address.
func (r *UserResolver) Email() string {
	return r.User.Email.String
}

func (r *UserResolver) Accounts(ctx context.Context) ([]*AccountResolver, error) {
	userID := r.User.ID

	accounts, err := r.Ctx.Repository.FindAccountsWithUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	result := make([]*AccountResolver, len(accounts))
	for idx, account := range accounts {
		result[idx] = &AccountResolver{
			Account: account,
			Ctx:     r.Ctx,
		}
	}
	return result, nil
}
