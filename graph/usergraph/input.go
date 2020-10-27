package usergraph

import "github.com/graph-gophers/graphql-go"

// UserArgs represents the args for getting user.
type UserArgs struct {
	ID graphql.ID
}

type CreateUserInput struct {
	Email string `validate:"email,required"`
	Name  string `validate:"required"`
}

type CreateUserArgs struct {
	Input CreateUserInput
}

type UpdateUserInput struct {
	ID   graphql.ID `validate:"required"`
	Name string     `validate:"required"`
}

type UpdateUserArgs struct {
	Input UpdateUserInput
}

type DeleteUserInput struct {
	ID graphql.ID `validate:"required"`
}

type DeleteUserArgs struct {
	Input DeleteUserInput
}
