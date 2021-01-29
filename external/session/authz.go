package session

import (
	"context"
	"errors"

	"github.com/alextanhongpin/pkg/contextkey"
	"github.com/google/uuid"
)

var userIDKey = contextkey.Key("user")

var ErrUnauthorized = errors.New("unauthorized")

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
