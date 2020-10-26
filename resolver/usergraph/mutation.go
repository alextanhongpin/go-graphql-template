package usergraph

import (
	"context"
	"errors"
	"fmt"

	"github.com/alextanhongpin/graphql-server-starter/model"
)

type Mutation struct {
	ctx *model.ResolverContext
}

func NewMutation(ctx *model.ResolverContext) *Mutation {
	return &Mutation{ctx: ctx}
}

type CreateUserInput struct {
	Email    string
	Password string
}

type CreateUserArgs struct {
	Input CreateUserInput
}

func (m *Mutation) CreateUser(ctx context.Context, args CreateUserArgs) (*Resolver, error) {
	fmt.Println("Mutation:CreateUser", args)
	return nil, errors.New("not implemented")
}
