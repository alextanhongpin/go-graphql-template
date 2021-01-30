package middleware

import (
	"net/http"

	"github.com/alextanhongpin/go-graphql-template/domain/account"
	"github.com/alextanhongpin/go-graphql-template/domain/user"
	"github.com/alextanhongpin/go-graphql-template/external/session"
)

func UserService(svc user.Service) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := session.WithUserService(r.Context(), svc)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AccountService(svc account.Service) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := session.WithAccountService(r.Context(), svc)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
