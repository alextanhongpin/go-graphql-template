package graph

import "github.com/graph-gophers/graphql-go"

// AccountArgs represents the args for getting account.
type AccountArgs struct {
	ID graphql.ID
}

type CreateAccountInput struct {
	UserID   graphql.ID
	Email    string
	Password string
}

type CreateAccountArgs struct {
	Input CreateAccountInput
}

type UpdateAccountInput struct {
	ID    graphql.ID
	Email string
}

type UpdateAccountArgs struct {
	Input UpdateAccountInput
}

type DeleteAccountInput struct {
	ID graphql.ID
}

type DeleteAccountArgs struct {
	Input DeleteAccountInput
}
