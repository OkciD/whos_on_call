package middleware

import (
	"net/http"

	loggerPkg "github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

func NewAccessLogMiddleware(logger loggerPkg.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// todo: response code, latency
			logger.WithRequest(r).Info("access log")

			next.ServeHTTP(w, r)
		})
	}
}
