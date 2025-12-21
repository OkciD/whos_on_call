package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/devicefeature"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type DeviceFeatureHandler struct {
	logger logger.Logger

	deviceFeatureUseCase devicefeature.UseCase
}

func New(mux *http.ServeMux, logger logger.Logger, deviceFeatureUseCase devicefeature.UseCase) *DeviceFeatureHandler {
	h := &DeviceFeatureHandler{
		logger: logger,

		deviceFeatureUseCase: deviceFeatureUseCase,
	}

	mux.Handle("/api/v1/device/{deviceid}/feature", handler.GenericHandler(logger, map[string]handler.Handler{
		http.MethodPost: h.Update,
	}))

	return h
}
