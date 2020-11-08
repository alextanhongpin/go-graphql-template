package middleware

import (
	"context"
	"net/http"

	"github.com/alextanhongpin/go-graphql-template/entity"
	"github.com/alextanhongpin/go-graphql-template/pkg/loader"
	"github.com/alextanhongpin/pkg/contextkey"
)

type DataLoaders struct {
	User loader.UserLoader
}

func NewDataLoaders(repo entity.Querier) *DataLoaders {
	return &DataLoaders{
		User: loader.NewUser(repo),
	}
}

var dataLoaderContext = contextkey.Key("dataloader")

func DataLoader(next http.Handler, repo entity.Querier) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := dataLoaderContext.WithValue(r.Context(), NewDataLoaders(repo))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ContextDataLoader(ctx context.Context) *DataLoaders {
	loader, _ := dataLoaderContext.Value(ctx).(*DataLoaders)
	return loader
}
