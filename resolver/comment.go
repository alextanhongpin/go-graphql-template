package resolver

import (
	"strings"

	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/alextanhongpin/graphql-server-starter/model"

	"github.com/graph-gophers/graphql-go"
)

// CommentResolver holds the comment entity to resolve.
type CommentResolver struct {
	Comment entity.Comment
	Ctx     *model.ResolverContext
}

// NewCommentResolver returns a new Comment resolver.
func NewCommentResolver(ctx *model.ResolverContext, comment entity.Comment) *CommentResolver {
	return &CommentResolver{
		Comment: comment,
		Ctx:     ctx,
	}
}

// ID returns the comment's id.
func (r *CommentResolver) ID() graphql.ID {
	return graphql.ID(r.Comment.ID.String())
}

// Provider returns the comment's provider.
func (r *CommentResolver) Provider() string {
	provider := string(r.Comment.Provider)
	return strings.ToUpper(provider)
}

// Email returns the comment's unique email.
func (r *CommentResolver) Email() string {
	return r.Comment.Email
}
