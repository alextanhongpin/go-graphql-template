package middleware

import (
	"net/http"

	"github.com/alextanhongpin/go-graphql-template/domain/user"
	"github.com/alextanhongpin/go-graphql-template/external/graph/loader"
	"github.com/alextanhongpin/go-graphql-template/external/session"
)

func DataLoader(usersvc user.Service) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// A new dataloader is created for each requests.
			ctx := session.WithDataLoader(r.Context(), loader.NewDataLoader(usersvc))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
