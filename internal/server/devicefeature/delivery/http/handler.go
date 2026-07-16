package http

import (
	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type DeviceFeatureHandler struct {
	logger logger.Logger

	deviceFeatureUseCase devicefeature.UseCase
}

func New(logger logger.Logger, deviceFeatureUseCase devicefeature.UseCase) DeviceFeatureHandler {
	return DeviceFeatureHandler{
		logger: logger,

		deviceFeatureUseCase: deviceFeatureUseCase,
	}
}
