package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	REQUEST_ID_CTX_KEY string = "requestId"
	REQUEST_ID_HEADER  string = "X-Request-Id"
)

func NewRequestIdMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqId := uuid.New().String()

			contextWithReqId := context.WithValue(r.Context(), REQUEST_ID_CTX_KEY, reqId)
			requestWithReqId := r.WithContext(contextWithReqId)

			w.Header().Add(REQUEST_ID_HEADER, reqId)

			next.ServeHTTP(w, requestWithReqId)
		})
	}
}

func GetRequestIdFromRequest(r *http.Request) string {
	reqId := r.Context().Value(REQUEST_ID_CTX_KEY).(string)
	if reqId == "" {
		return "undef"
	}

	return reqId
}
