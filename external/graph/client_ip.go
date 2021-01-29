package graph

import (
	"context"
	"net"
	"net/http"
	"strings"

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

// The implementation is similar to gin-gonic's .ClientIP method.
func clientIPFromRequest(r *http.Request) string {
	clientIP := r.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err != nil {
		return ip
	}
	return ""
}

func ClientIPProvider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := clientIPFromRequest(r)
		ctx := r.Context()
		ctx = WithClientIP(ctx, clientIP)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
