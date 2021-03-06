package resolver

import (
	"context"
	"strings"

	"github.com/alextanhongpin/go-graphql-template/domain/account"
	"github.com/alextanhongpin/go-graphql-template/external/session"

	"github.com/graph-gophers/graphql-go"
)

// AccountResolver holds the account entity to resolve.
type AccountResolver struct {
	account account.Account
}

// NewAccountResolver returns a new Account resolver.
func NewAccountResolver(acc account.Account) *AccountResolver {
	return &AccountResolver{
		account: acc,
	}
}

// ID returns the account's id.
func (r *AccountResolver) ID() graphql.ID {
	return graphql.ID(r.account.ID.String())
}

// Provider returns the account's provider.
func (r *AccountResolver) Provider() string {
	provider := string(r.account.Provider)
	return strings.ToUpper(provider)
}

// Email returns the account's unique email.
func (r *AccountResolver) Email() string {
	return r.account.Email
}

// User returns the account's user.
func (r *AccountResolver) User(ctx context.Context) (*PartialUserResolver, error) {
	userID := r.account.UserID
	loader := session.DataLoader(ctx)
	user, err := loader.User.Load(ctx, userID.String())
	if err != nil {
		return nil, err
	}

	return NewPartialUserResolver(user), nil
}
