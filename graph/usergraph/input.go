package usergraph

import "github.com/graph-gophers/graphql-go"

// UserArgs represents the args for getting user.
type UserArgs struct {
	ID graphql.ID
}

type CreateUserInput struct {
	Email    string
	Name     string
	Password string
}

type CreateUserArgs struct {
	Input CreateUserInput
}

type UpdateUserInput struct {
	ID   graphql.ID
	Name string
}

type UpdateUserArgs struct {
	Input UpdateUserInput
}

type DeleteUserInput struct {
	ID graphql.ID
}

type DeleteUserArgs struct {
	Input DeleteUserInput
}
