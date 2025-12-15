package handler

import (
	"net/http"

	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
	loggerPkg "github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type ResponseWriter func(w http.ResponseWriter) error

type Handler func(r *http.Request) (ResponseWriter, error)

func GenericHandler(logger loggerPkg.Logger, handlers map[string]Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler, ok := handlers[r.Method]
		if !ok {
			logger.WithFields(loggerPkg.Fields{
				"method":     r.Method,
				"requestURI": r.RequestURI,
			}).Error("handler not found")

			RespondJSONError(w, appErrors.ErrNotImplemented)

			return
		}

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
