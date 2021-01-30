package session

import (
	"context"
	"fmt"

	"github.com/alextanhongpin/go-graphql-template/domain/account"
	"github.com/alextanhongpin/go-graphql-template/domain/user"
	"github.com/alextanhongpin/pkg/contextkey"
)

var (
	usersvcKey    = contextkey.Key("usersvc")
	accountsvcKey = contextkey.Key("accountsvc")
)

func UserService(ctx context.Context) (user.Service, error) {
	svc, ok := usersvcKey.Value(ctx).(user.Service)
	if !ok {
		return nil, fmt.Errorf("missing usersvc context")
	}
	return svc, nil
}

func WithUserService(ctx context.Context, svc user.Service) context.Context {
	return context.WithValue(ctx, usersvcKey, svc)
}

func AccountService(ctx context.Context) (account.Service, error) {
	svc, ok := accountsvcKey.Value(ctx).(account.Service)
	if !ok {
		return nil, fmt.Errorf("missing accountsvc context")
	}
	return svc, nil
}

func WithAccountService(ctx context.Context, svc account.Service) context.Context {
	return context.WithValue(ctx, accountsvcKey, svc)
}
