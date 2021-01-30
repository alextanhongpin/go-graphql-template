package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/alextanhongpin/go-graphql-template/external/security"
	"github.com/alextanhongpin/go-graphql-template/external/session"
	"github.com/google/uuid"
)

var ErrInvalidAuthHeader = errors.New("Authorization header is invalid")

func Authz(authz security.Authorizer) Middleware {
	return func(next http.Handler) http.Handler {
		authorize := func(auth string) (*security.Claims, error) {
			if auth == "" {
				return nil, nil
			}

			paths := strings.Fields(auth)
			if len(paths) != 2 {
				return nil, ErrInvalidAuthHeader
			}

			bearer, token := paths[0], paths[1]
			if isBearerValid := strings.EqualFold("bearer", bearer); !isBearerValid {
				return nil, ErrInvalidAuthHeader
			}

			claims, err := authz.Verify(token)
			if err != nil {
				return nil, fmt.Errorf("invalid token: %w", err)
			}
			return claims, nil
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := authorize(r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			ctx := r.Context()
			if claims != nil {
				userID, err := uuid.Parse(claims.StandardClaims.Subject)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				ctx = session.WithUserID(ctx, userID)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
