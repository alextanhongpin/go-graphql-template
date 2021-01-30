package account

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/internal"
	"github.com/google/uuid"
)

type Service interface {
	FindAccount(ctx context.Context, id uuid.UUID) (Account, error)
	FindAccountsWithUserID(ctx context.Context, userID uuid.UUID) ([]Account, error)

	CreateAccount(ctx context.Context, dto CreateAccountDto) (Account, error)
	UpdateAccount(ctx context.Context, dto UpdateAccountDto) (Account, error)
	DeleteAccount(ctx context.Context, accountID uuid.UUID) (Account, error)
}

var noop *Account

func init() {
	noop = &Account{}
}

type AccountService struct {
	repo      Repository
	validator internal.Validator
}

func NewService(r Repository, v internal.Validator) *AccountService {
	return &AccountService{
		repo:      r,
		validator: v,
	}
}

func (svc *AccountService) FindAccount(ctx context.Context, accountID uuid.UUID) (Account, error) {
	a, err := svc.repo.FindAccount(ctx, accountID)
	if err != nil {
		return *noop, err
	}

	return ToAccount(a), nil
}

func (svc *AccountService) FindAccountsWithUserID(ctx context.Context, userID uuid.UUID) ([]Account, error) {
	accounts, err := svc.repo.FindAccountsWithUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return Accounts(accounts).ToAccounts(), nil
}

func (svc *AccountService) CreateAccount(ctx context.Context, dto CreateAccountDto) (Account, error) {
	if err := svc.validator.Validate(ctx, &dto); err != nil {
		return *noop, err
	}

	a, err := svc.repo.CreateAccount(ctx, dto.ToRepoCreateAccount())
	if err != nil {
		return *noop, err
	}
	return ToAccount(a), nil
}

func (svc *AccountService) UpdateAccount(ctx context.Context, dto UpdateAccountDto) (Account, error) {
	if err := svc.validator.Validate(ctx, &dto); err != nil {
		return *noop, err
	}

	a, err := svc.repo.UpdateAccount(ctx, dto.ToRepoUpdateAccount())
	if err != nil {
		return *noop, err
	}

	return ToAccount(a), nil
}

func (svc *AccountService) DeleteAccount(ctx context.Context, accountID uuid.UUID) (Account, error) {
	a, err := svc.repo.DeleteAccount(ctx, accountID)
	if err != nil {
		return *noop, err
	}

	return ToAccount(a), nil
}
