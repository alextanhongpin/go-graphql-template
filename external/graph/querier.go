package graph

import (
	"context"
	"errors"
	"net/http"

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

func BuildQuerierProvider(q entity.Querier) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := WithQuerier(r.Context(), q)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
