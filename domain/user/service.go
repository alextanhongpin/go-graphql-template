package user

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/internal"
	"github.com/google/uuid"
)

type Service interface {
	FindUser(ctx context.Context, id uuid.UUID) (User, error)
	FindUsersWithIDs(ctx context.Context, ids []uuid.UUID) ([]User, error)

	CreateUser(ctx context.Context, dto CreateUserDto) (User, error)
	UpdateUser(ctx context.Context, dto UpdateUserDto) (User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) (User, error)
}

var nouser *User

func init() {
	nouser = &User{}
}

type UserService struct {
	repo      Repository
	validator internal.Validator
}

func NewService(r Repository, v internal.Validator) *UserService {
	return &UserService{
		repo:      r,
		validator: v,
	}
}

func (svc *UserService) FindUsersWithIDs(ctx context.Context, ids []uuid.UUID) ([]User, error) {
	if len(ids) == 0 {
		return make([]User, 0), nil
	}

	users, err := svc.repo.FindUsersWithIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return Users(users).ToUsers(), nil
}

func (svc *UserService) FindUser(ctx context.Context, userID uuid.UUID) (User, error) {
	u, err := svc.repo.FindUser(ctx, userID)
	if err != nil {
		return *nouser, err
	}

	return ToUser(u), nil
}

func (svc *UserService) CreateUser(ctx context.Context, dto CreateUserDto) (User, error) {
	if err := svc.validator.Validate(ctx, &dto); err != nil {
		return *nouser, err
	}

	u, err := svc.repo.CreateUser(ctx, dto.ToRepoCreateUser())
	if err != nil {
		return *nouser, err
	}

	return ToUser(u), nil
}

func (svc *UserService) UpdateUser(ctx context.Context, dto UpdateUserDto) (User, error) {
	if err := svc.validator.Validate(ctx, &dto); err != nil {
		return *nouser, err
	}

	u, err := svc.repo.UpdateUser(ctx, dto.ToRepoUpdateUser())
	if err != nil {
		return *nouser, err
	}

	return ToUser(u), nil
}

func (svc *UserService) DeleteUser(ctx context.Context, userID uuid.UUID) (User, error) {
	u, err := svc.repo.DeleteUser(ctx, userID)
	if err != nil {
		return *nouser, err
	}

	return ToUser(u), nil
}
