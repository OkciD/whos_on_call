package handler

import (
	"net/http"

	loggerPkg "github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type ResponseWriter func(w http.ResponseWriter) error

type Handler func(r *http.Request) (ResponseWriter, error)

func GenericHandler(logger loggerPkg.Logger, handler Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logger.WithContext(r.Context())

		responseWriter, err := handler(r)
		if err != nil {
			logger.WithError(err).Error("error returned from handler")
			RespondJSONError(w, err)
			return
		}

		err = responseWriter(w)
		if err != nil {
			logger.WithError(err).Error("error returned from response writer")
			RespondJSONError(w, err)
			return
		}
	})
}
