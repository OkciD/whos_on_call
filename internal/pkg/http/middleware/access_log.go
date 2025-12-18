package middleware

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/context"
	loggerPkg "github.com/OkciD/whos_on_call/internal/pkg/logger"
)

func NewAccessLogMiddleware(logger loggerPkg.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				ip     = r.RemoteAddr
				method = r.Method
				url    = r.URL.String()
				reqId  = context.GetRequestId(r.Context())
			)

			logger.WithFields(loggerPkg.Fields{
				"ip":     ip,
				"method": method,
				"url":    url,
				"reqId":  reqId,
			}).Info("access log")

			next.ServeHTTP(w, r)
		})
	}
}
