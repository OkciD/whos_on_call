package http

import (
	"github.com/OkciD/whos_on_call/internal/server/device"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type DeviceHandler struct {
	logger logger.Logger

	deviceUseCase device.UseCase
}

func New(logger logger.Logger, deviceUseCase device.UseCase) DeviceHandler {
	return DeviceHandler{
		logger: logger,

		deviceUseCase: deviceUseCase,
	}
}
