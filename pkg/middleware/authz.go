package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/alextanhongpin/pkg/contextkey"
	"github.com/alextanhongpin/pkg/gojwt"
)

type signer interface {
	Verify(token string) (*gojwt.Claims, error)
}

var (
	ErrEmptyAuthHeader   = errors.New("Authorization header is not provided")
	ErrInvalidAuthHeader = errors.New("Authorization header is invalid")
	ErrUnauthorized      = errors.New("unauthorized")
)

var (
	UserContext = contextkey.Key("user")
)

func ContextUserID(ctx context.Context) (uuid.UUID, error) {
	id, ok := UserContext.Value(ctx).(string)
	if !ok {
		return uuid.UUID{}, ErrUnauthorized
	}
	return uuid.Parse(id)
}

func Authz(next http.Handler, signer signer) http.Handler {
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

		claims, err := signer.Verify(token)
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
			ctx = UserContext.WithValue(ctx, claims.StandardClaims.Subject)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
