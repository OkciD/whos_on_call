package middleware

import (
	"net/http"

	"github.com/google/uuid"

	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

const ReqIDHeader string = "X-Request-Id"

func NewRequestIDMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := uuid.New().String()

			contextWithReqID := appContext.StoreRequestID(r.Context(), reqID)
			contextWithLoggerReqID := logger.AddFieldsToContext(contextWithReqID, logger.Fields{
				"reqid": reqID,
			})

			requestWithReqID := r.WithContext(contextWithLoggerReqID)

			w.Header().Add(ReqIDHeader, reqID)

			next.ServeHTTP(w, requestWithReqID)
		})
	}
}
