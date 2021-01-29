package graph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/alextanhongpin/pkg/contextkey"
	"github.com/alextanhongpin/pkg/gojwt"
	"github.com/google/uuid"
)

var userIDKey = contextkey.Key("user")

var (
	ErrEmptyAuthHeader   = errors.New("Authorization header is not provided")
	ErrInvalidAuthHeader = errors.New("Authorization header is invalid")
	ErrUnauthorized      = errors.New("unauthorized")
)

func UserID(ctx context.Context) (uuid.UUID, error) {
	id, ok := userIDKey.Value(ctx).(uuid.UUID)
	if !ok {
		return id, ErrUnauthorized
	}
	return id, nil
}

func WithUserID(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

type Authorizer interface {
	Verify(token string) (*gojwt.Claims, error)
}

func BuildAuthz(authz Authorizer) Middleware {
	return func(next http.Handler) http.Handler {
		authorize := func(auth string) (*gojwt.Claims, error) {
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
				ctx = WithUserID(ctx, userID)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
