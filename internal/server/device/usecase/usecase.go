package usecase

import (
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	"github.com/OkciD/whos_on_call/internal/server/device"
)

type UseCase struct {
	logger logger.Logger

	deviceRepo device.Repository
}

func New(logger logger.Logger, deviceRepo device.Repository) device.UseCase {
	return &UseCase{
		logger: logger,

		deviceRepo: deviceRepo,
	}
}
