package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func LimiterMiddleware(next http.Handler) http.Handler {
	limit := rate.NewLimiter(1, 2)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limit.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`rate limit exceeded.`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
