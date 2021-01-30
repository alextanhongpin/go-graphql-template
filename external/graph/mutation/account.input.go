package mutation

import (
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"

	"github.com/alextanhongpin/go-graphql-template/domain/account"
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
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

func (i CreateAccountInput) ToServiceCreateAccount() (account.CreateAccountDto, error) {
	var dto account.CreateAccountDto

	userID, err := uuid.Parse(string(i.UserID))
	if err != nil {
		return dto, err
	}

	return account.CreateAccountDto{
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

func (i UpdateAccountInput) ToServiceUpdateAccount() (account.UpdateAccountDto, error) {
	var dto account.UpdateAccountDto
	id, err := uuid.Parse(string(i.ID))
	if err != nil {
		return dto, err
	}

	return account.UpdateAccountDto{
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

func (i DeleteAccountInput) ToServiceDeleteAccount() (uuid.UUID, error) {
	return uuid.Parse(string(i.ID))
}
