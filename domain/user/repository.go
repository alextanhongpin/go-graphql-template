package user

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/google/uuid"
)

type Repository interface {
	FindUser(ctx context.Context, id uuid.UUID) (entity.User, error)
	FindUsersWithIDs(ctx context.Context, ids []uuid.UUID) ([]entity.User, error)

	CreateUser(ctx context.Context, params entity.CreateUserParams) (entity.User, error)
	UpdateUser(ctx context.Context, params entity.UpdateUserParams) (entity.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (entity.User, error)
}
