package graph

import (
	"context"
	"net/http"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/pkg/loader"
	"github.com/alextanhongpin/pkg/contextkey"
)

var dataLoaderKey = contextkey.Key("dataloader")

func DataLoader(ctx context.Context) *DataLoaders {
	dl, _ := dataLoaderKey.Value(ctx).(*DataLoaders)
	return dl
}

func WithDataLoader(ctx context.Context, dl *DataLoaders) context.Context {
	return context.WithValue(ctx, dataLoaderKey, dl)
}

func BuildDataLoaderProvider(dl *DataLoaders) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := WithDataLoader(r.Context(), dl)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

type DataLoaders struct {
	User loader.UserLoader
}

func NewDataLoaders(repo entity.Querier) *DataLoaders {
	return &DataLoaders{
		User: loader.NewUser(repo),
	}
}
