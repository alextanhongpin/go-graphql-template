package mutation

import (
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
)

// We can also place inputs in a folder called inputs.

type CreateAccountArgs struct {
	Input CreateAccountInput
}

type CreateAccountInput struct {
	UserID   graphql.ID
	Email    string
	Password string
}

func (i CreateAccountInput) ToRepoCreateAccount() (entity.CreateAccountParams, error) {
	var p entity.CreateAccountParams

	userID, err := uuid.Parse(string(i.UserID))
	if err != nil {
		return p, err
	}

	return entity.CreateAccountParams{
		Uid:      i.Email,
		Provider: entity.ProviderEmail,
		Token:    i.Password,
		Email:    i.Email,
		UserID:   userID,
	}, nil
}

type UpdateAccountArgs struct {
	Input UpdateAccountInput
}

type UpdateAccountInput struct {
	ID    graphql.ID
	Email string
}

func (i UpdateAccountInput) ToRepoUpdateAccount() (entity.UpdateAccountParams, error) {
	var p entity.UpdateAccountParams
	id, err := uuid.Parse(string(i.ID))
	if err != nil {
		return p, err
	}

	return entity.UpdateAccountParams{
		ID:    id,
		Email: i.Email,
	}, nil
}

type DeleteAccountArgs struct {
	Input DeleteAccountInput
}

type DeleteAccountInput struct {
	ID graphql.ID
}

func (i DeleteAccountInput) ToRepoDeleteAccount() (uuid.UUID, error) {
	return uuid.Parse(string(i.ID))
}
