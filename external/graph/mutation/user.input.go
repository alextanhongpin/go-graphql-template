package mutation

import (
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"

	"github.com/alextanhongpin/go-graphql-template/domain/user"
)

type CreateUserArgs struct {
	Input CreateUserInput
}

type CreateUserInput struct {
	Email string
	Name  string
}

func (i CreateUserInput) ToServiceCreateUser() user.CreateUserDto {
	return user.CreateUserDto{
		Name:              i.Name,
		Email:             i.Email,
		PreferredUsername: i.Name,
	}
}

type UpdateUserArgs struct {
	Input UpdateUserInput
}

type UpdateUserInput struct {
	ID   graphql.ID
	Name string
}

func (i UpdateUserInput) ToServiceUpdateUser() (user.UpdateUserDto, error) {
	var p user.UpdateUserDto
	userID, err := uuid.Parse(string(i.ID))
	if err != nil {
		return p, err
	}

	return user.UpdateUserDto{
		ID:   userID,
		Name: i.Name,
	}, nil
}

type DeleteUserArgs struct {
	Input DeleteUserInput
}

type DeleteUserInput struct {
	ID graphql.ID
}

func (i DeleteUserInput) ToServiceDeleteUser() (uuid.UUID, error) {
	return uuid.Parse(string(i.ID))
}
