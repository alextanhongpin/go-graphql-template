package domain

import (
	"github.com/alextanhongpin/go-graphql-template/domain/account"
	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/domain/user"
	"github.com/google/wire"
)

var AccountSet = wire.NewSet(
	account.NewService,
	wire.FieldsOf(new(*Repository), "Account"),
	wire.Bind(new(account.Service), new(*account.AccountService)),
)

var UserSet = wire.NewSet(
	user.NewService,
	wire.FieldsOf(new(*Repository), "User"),
	wire.Bind(new(user.Service), new(*user.UserService)),
)

var Set = wire.NewSet(
	NewRepository,
	AccountSet,
	UserSet,
)

type Repository struct {
	User    user.Repository
	Account account.Repository
}

func NewRepository(r entity.Querier) *Repository {
	return &Repository{
		Account: r,
		User:    r,
	}
}
