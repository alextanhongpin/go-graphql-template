package middleware

import (
	"net/http"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/external/session"
	"github.com/alextanhongpin/go-graphql-template/pkg/loader"
)

func DataLoader(q entity.Querier) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// A new dataloader is created for each requests.
			ctx := session.WithDataLoader(r.Context(), loader.NewDataLoader(q))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
