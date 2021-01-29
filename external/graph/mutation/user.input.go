package mutation

import (
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/domain/model"
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
)

type CreateUserArgs struct {
	Input CreateUserInput
}

type CreateUserInput struct {
	Email string `validate:"email,required"`
	Name  string `validate:"required"`
}

func (c CreateUserInput) ToRepoCreateUser() entity.CreateUserParams {
	return entity.CreateUserParams{
		Name:              c.Name,
		Email:             model.NewNullString(c.Email),
		PreferredUsername: c.Name,
	}
}

type UpdateUserArgs struct {
	Input UpdateUserInput
}

type UpdateUserInput struct {
	ID   graphql.ID `validate:"required"`
	Name string     `validate:"required"`
}

func (u UpdateUserInput) ToRepoUpdateUser() (entity.UpdateUserParams, error) {
	var p entity.UpdateUserParams
	userID, err := uuid.Parse(string(u.ID))
	if err != nil {
		return p, err
	}

	return entity.UpdateUserParams{
		ID:   userID,
		Name: u.Name,
	}, nil
}

type DeleteUserArgs struct {
	Input DeleteUserInput
}

type DeleteUserInput struct {
	ID graphql.ID `validate:"required"`
}

func (i DeleteUserInput) ToRepoDeleteUser() (uuid.UUID, error) {
	return uuid.Parse(string(i.ID))
}
