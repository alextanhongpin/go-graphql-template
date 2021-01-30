package query

import (
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
)

// UserArgs represents the args for getting user.
type UserArgs struct {
	ID graphql.ID
}

func (u UserArgs) ToServiceFindUser() (uuid.UUID, error) {
	return uuid.Parse(string(u.ID))
}
