package apiserver

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/shared/errors/mapper"
	loggerPkg "github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

func NewResponseErrorHandler(logger loggerPkg.Logger) func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		logger := logger.WithContext(r.Context()).WithError(err)

		logger.Error("response error")

		respondError(w, mapper.ErrorToResp(err))
	}
}
