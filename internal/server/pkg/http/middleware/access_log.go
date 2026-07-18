package middleware

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/server/pkg/context"
	loggerPkg "github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

func NewAccessLogMiddleware(logger loggerPkg.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.WithRequest(r).WithField(
				"reqId", context.GetRequestId(r.Context()),
			).Info("access log")

			next.ServeHTTP(w, r)
		})
	}
}
