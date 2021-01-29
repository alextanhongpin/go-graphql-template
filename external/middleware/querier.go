package middleware

import (
	"net/http"

	"github.com/alextanhongpin/go-graphql-template/domain/entity"
	"github.com/alextanhongpin/go-graphql-template/external/session"
)

func Querier(q entity.Querier) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := session.WithQuerier(r.Context(), q)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
