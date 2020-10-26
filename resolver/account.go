package resolver

import (
	"context"
	"strings"

	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/model"

	"github.com/graph-gophers/graphql-go"
)

// AccountResolver holds the account entity to resolve.
type AccountResolver struct {
	Account entity.Account
	Ctx     *model.ResolverContext
}

// NewAccountResolver returns a new Account resolver.
func NewAccountResolver(ctx *model.ResolverContext, account entity.Account) *AccountResolver {
	return &AccountResolver{
		Account: account,
		Ctx:     ctx,
	}
}

// ID returns the account's id.
func (r *AccountResolver) ID() graphql.ID {
	return graphql.ID(r.Account.ID.String())
}

// Provider returns the account's provider.
func (r *AccountResolver) Provider() string {
	provider := string(r.Account.Provider)
	return strings.ToUpper(provider)
}

// Email returns the account's unique email.
func (r *AccountResolver) Email() string {
	return r.Account.Email
}

// User returns the account's user.
func (r *AccountResolver) User(ctx context.Context) (*UserResolver, error) {
	userID := r.Account.UserID
	user, err := r.Ctx.Repository.FindUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &UserResolver{
		User: user,
		Ctx:  r.Ctx,
	}, nil
}
