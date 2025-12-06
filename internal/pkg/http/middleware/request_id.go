package middleware

import (
	"context"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	"github.com/google/uuid"
)

const REQUEST_ID_HEADER string = "X-Request-Id"

type requestIdCtxKey = struct{}

func NewRequestIdMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqId := uuid.New().String()

			contextWithReqId := context.WithValue(r.Context(), requestIdCtxKey{}, reqId)
			contextWithLoggerReqId := logger.AddFieldsToContext(contextWithReqId, logger.Fields{
				"reqid": reqId,
			})

			requestWithReqId := r.WithContext(contextWithLoggerReqId)

			w.Header().Add(REQUEST_ID_HEADER, reqId)

			next.ServeHTTP(w, requestWithReqId)
		})
	}
}

func GetRequestIdFromRequest(r *http.Request) string {
	reqId, ok := r.Context().Value(requestIdCtxKey{}).(string)
	if !ok {
		return "undef"
	}

	return reqId
}
