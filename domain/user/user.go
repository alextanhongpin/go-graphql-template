package user

import (
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID
	Name              string
	Email             string
	PreferredUsername string
}

func ToUser(e entity.User) User {
	return User{
		ID:                e.ID,
		Name:              e.Name,
		Email:             e.Email.String,
		PreferredUsername: e.PreferredUsername,
	}
}

type Users []entity.User

func (u Users) ToUsers() []User {
	users := make([]User, len(u))
	for i, usr := range u {
		users[i] = ToUser(usr)
	}
	return users
}
