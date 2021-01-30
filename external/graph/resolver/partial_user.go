package resolver

import (
	"github.com/alextanhongpin/go-graphql-template/domain/user"

	"github.com/graph-gophers/graphql-go"
)

// PartialUserResolver returns the partial field for users, excluding
// associations to avoid circular reference.
// For example, for a UserResolver and CommentResolver, UserResolver should
// return CommentResolver, but CommentResolver should return
// PartialUserResolver.
type PartialUserResolver struct {
	u user.User
}

// NewPartialUserResolver returns a new PartialUser resolver.
func NewPartialUserResolver(u user.User) *PartialUserResolver {
	return &PartialUserResolver{u}
}

// ID returns the user's id.
func (r *PartialUserResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID.String())
}

// Name returns the user's name.
func (r *PartialUserResolver) Name() string {
	return r.u.Name
}

// Email returns the user's unique email address.
func (r *PartialUserResolver) Email() string {
	return r.u.Email
}
