package middleware

import (
	"net/http"

	"runtime/debug"

	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
	loggerPkg "github.com/OkciD/whos_on_call/internal/pkg/logger"
)

func NewRecoveryMiddleware(logger loggerPkg.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.WithFields(loggerPkg.Fields{
						"panic": err,
						"stack": string(debug.Stack()),
					}).Error("panic occurred")

					handler.RespondJSONError(w, appErrors.ErrInternal)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
