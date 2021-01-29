package resolver_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/domain/model"
)

type userRepo struct {
	entity.Querier
	accounts    []entity.Account
	accountsErr error
}

func (u *userRepo) FindAccountsWithUserID(ctx context.Context, userID uuid.UUID) ([]entity.Account, error) {
	return u.accounts, u.accountsErr
}

func TestUserResolver(t *testing.T) {
	frozenTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)

	repo := &userRepo{
		accounts: []entity.Account{
			entity.Account{
				ID:        uuid.New(),
				Uid:       "random uid",
				Provider:  entity.ProviderEmail,
				Token:     "random token",
				UserID:    uuid.New(),
				Email:     "john.doe@mail.com",
				Data:      json.RawMessage{},
				CreatedAt: frozenTime,
				DeletedAt: sql.NullTime{
					Time:  frozenTime,
					Valid: true,
				},
			},
		},
		accountsErr: nil,
	}
	assert := assert.New(t)
	user := entity.User{
		ID:    uuid.New(),
		Name:  "John Doe",
		Email: model.NewNullString("john.doe@mail.com"),
	}

	userResolver := graph.NewUserResolver(graph.NewContext(repo), user)
	assert.Equal(user.ID.String(), string(userResolver.ID()))
	assert.Equal(user.Name, userResolver.Name())
	assert.Equal(user.Email.String, userResolver.Email())
	accounts, err := userResolver.Accounts(context.TODO())
	assert.Nil(err)
	assert.True(len(accounts) == 1)
}
