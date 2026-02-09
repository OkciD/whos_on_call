package middleware

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/google/uuid"
)

const REQUEST_ID_HEADER string = "X-Request-Id"

func NewRequestIdMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqId := uuid.New().String()

			contextWithReqId := appContext.StoreRequestId(r.Context(), reqId)
			contextWithLoggerReqId := logger.AddFieldsToContext(contextWithReqId, logger.Fields{
				"reqid": reqId,
			})

			requestWithReqId := r.WithContext(contextWithLoggerReqId)

			w.Header().Add(REQUEST_ID_HEADER, reqId)

			next.ServeHTTP(w, requestWithReqId)
		})
	}
}
