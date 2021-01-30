package account

import (
	"github.com/google/uuid"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
)

type Account struct {
	ID       uuid.UUID       `json:"id"`
	Provider entity.Provider `json:"provider"`
	UserID   uuid.UUID       `json:"user_id"`
	Email    string          `json:"email"`
}

func ToAccount(acc entity.Account) Account {
	return Account{
		ID:       acc.ID,
		Provider: acc.Provider,
		UserID:   acc.UserID,
		Email:    acc.Email,
	}
}

type Accounts []entity.Account

func (a Accounts) ToAccounts() []Account {
	accounts := make([]Account, len(a))
	for i, acc := range a {
		accounts[i] = ToAccount(acc)
	}
	return accounts
}
