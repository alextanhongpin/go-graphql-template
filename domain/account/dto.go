package account

import (
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/google/uuid"
)

type CreateAccountDto struct {
	Uid      string          `json:"uid" validate:"required"`
	Provider entity.Provider `json:"provider" validate:"required"`
	Token    string          `json:"token" validate:"required"`
	UserID   uuid.UUID       `json:"userID" validate:"required"`
	Email    string          `json:"email" validate:"email,required"`
}

func (dto CreateAccountDto) ToRepoCreateAccount() entity.CreateAccountParams {
	return entity.CreateAccountParams(dto)
}

type UpdateAccountDto struct {
	Email string    `json:"email" validate:"email,required"`
	ID    uuid.UUID `json:"id" validate:"required"`
}

func (dto UpdateAccountDto) ToRepoUpdateAccount() entity.UpdateAccountParams {
	return entity.UpdateAccountParams(dto)
}
