package usecase

import (
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	"github.com/OkciD/whos_on_call/internal/server/device"
	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
)

type UseCase struct {
	logger logger.Logger

	deviceRepo        device.Repository
	deviceFeatureRepo devicefeature.Repository
}

func New(logger logger.Logger, deviceRepo device.Repository, deviceFeatureRepo devicefeature.Repository) devicefeature.UseCase {
	return &UseCase{
		logger: logger,

		deviceRepo:        deviceRepo,
		deviceFeatureRepo: deviceFeatureRepo,
	}
}
