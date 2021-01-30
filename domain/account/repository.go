package account

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/google/uuid"
)

type Repository interface {
	FindAccount(ctx context.Context, id uuid.UUID) (entity.Account, error)
	FindAccountsWithUserID(ctx context.Context, userID uuid.UUID) ([]entity.Account, error)

	CreateAccount(ctx context.Context, arg entity.CreateAccountParams) (entity.Account, error)

	UpdateAccount(ctx context.Context, arg entity.UpdateAccountParams) (entity.Account, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) (entity.Account, error)
}
