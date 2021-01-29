package resolver

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/external/session"

	"github.com/graph-gophers/graphql-go"
)

// UserResolver holds the user entity to resolve.
type UserResolver struct {
	user entity.User
}

// NewUserResolver returns a new User resolver.
func NewUserResolver(user entity.User) *UserResolver {
	return &UserResolver{user: user}
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

// Owner returns true if the authorized user owns this profile.
func (r *UserResolver) Owner(ctx context.Context) bool {
	userID, err := session.UserID(ctx)
	if err != nil {
		return false
	}
	return r.user.ID == userID
}

func (r *UserResolver) Accounts(ctx context.Context) ([]*AccountResolver, error) {
	userID := r.user.ID

	q, err := session.Querier(ctx)
	if err != nil {
		return nil, err
	}

	accounts, err := q.FindAccountsWithUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]*AccountResolver, len(accounts))
	for idx, account := range accounts {
		result[idx] = NewAccountResolver(account)
	}

	return result, nil
}
