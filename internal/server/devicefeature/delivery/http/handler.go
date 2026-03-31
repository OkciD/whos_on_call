package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
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

	mux.Handle("POST /api/v1/device/{deviceid}/feature", handler.GenericHandler(logger, h.Update))

	return h
}
