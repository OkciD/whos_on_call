package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func NewAccessLogMiddleware(logger *logrus.Entry) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				ip     = r.RemoteAddr
				method = r.Method
				url    = r.URL.String()
				reqId  = GetRequestIdFromRequest(r)
			)

			logger.WithFields(logrus.Fields{
				"ip":     ip,
				"method": method,
				"url":    url,
				"reqId":  reqId,
			}).Info("access log")

			next.ServeHTTP(w, r)
		})
	}
}
