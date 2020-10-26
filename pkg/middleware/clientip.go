package middleware

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/alextanhongpin/pkg/contextkey"
)

var clientIPContext = contextkey.Key("clientip")

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

func ClientIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := clientIPFromRequest(r)
		ctx := r.Context()
		ctx = clientIPContext.WithValue(ctx, clientIP)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ContextClientIP(ctx context.Context) string {
	clientIP, _ := clientIPContext.Value(ctx).(string)
	return clientIP
}
