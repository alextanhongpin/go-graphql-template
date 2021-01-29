package resolver

import (
	"github.com/alextanhongpin/go-graphql-template/domain/entity"

	"github.com/graph-gophers/graphql-go"
)

// PartialUserResolver returns the partial field for users, excluding
// associations to avoid circular reference.
// For example, for a UserResolver and CommentResolver, UserResolver should
// return CommentResolver, but CommentResolver should return
// PartialUserResolver.
type PartialUserResolver struct {
	user entity.User
}

// NewPartialUserResolver returns a new PartialUser resolver.
func NewPartialUserResolver(user entity.User) *PartialUserResolver {
	return &PartialUserResolver{
		user: user,
	}
}

// ID returns the user's id.
func (r *PartialUserResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID.String())
}

// Name returns the user's name.
func (r *PartialUserResolver) Name() string {
	return r.user.Name
}

// Email returns the user's unique email address.
func (r *PartialUserResolver) Email() string {
	return r.user.Email.String
}
