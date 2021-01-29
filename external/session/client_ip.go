package session

import (
	"context"

	"github.com/alextanhongpin/pkg/contextkey"
)

var clientIPKey = contextkey.Key("client-ip")

func ClientIP(ctx context.Context) string {
	clientIP, _ := clientIPKey.Value(ctx).(string)
	return clientIP
}

func WithClientIP(ctx context.Context, clientIP string) context.Context {
	return context.WithValue(ctx, clientIPKey, clientIP)
}
