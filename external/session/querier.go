package session

import (
	"context"
	"errors"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/pkg/contextkey"
)

var querierKey = contextkey.Key("user")

func Querier(ctx context.Context) (entity.Querier, error) {
	q, ok := querierKey.Value(ctx).(entity.Querier)
	if !ok {
		return nil, errors.New("ctx querier not provided")
	}
	return q, nil
}

func WithQuerier(ctx context.Context, q entity.Querier) context.Context {
	return context.WithValue(ctx, querierKey, q)
}
