package user

import (
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/domain/model"
	"github.com/google/uuid"
)

type CreateUserDto struct {
	Name              string `json:"name" validate:"required"`
	Email             string `json:"email" validate:"email,required"`
	PreferredUsername string `json:"preferredUsername" validate:"required"`
}

// Why dto vs c? it's easier when you rename if you name it by functionality rather than class name.
func (dto CreateUserDto) ToRepoCreateUser() entity.CreateUserParams {
	return entity.CreateUserParams{
		Name:              dto.Name,
		Email:             model.NewNullString(dto.Email),
		PreferredUsername: dto.PreferredUsername,
	}
}

type UpdateUserDto struct {
	ID   uuid.UUID `json:"id" validate:"required"`
	Name string    `json:"name" validate:"required,min=3"`
}

func (dto UpdateUserDto) ToRepoUpdateUser() entity.UpdateUserParams {
	return entity.UpdateUserParams{
		ID:   dto.ID,
		Name: dto.Name,
	}
}
