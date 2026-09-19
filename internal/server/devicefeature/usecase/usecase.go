package usecase

import (
	"github.com/OkciD/whos_on_call/internal/server/device"
	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/eventbus"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type UseCase struct {
	logger logger.Logger
	eb     eventbus.EventBus

	deviceRepo        device.Repository
	deviceFeatureRepo devicefeature.Repository
}

func New(
	logger logger.Logger,
	eventBus eventbus.EventBus,
	deviceRepo device.Repository,
	deviceFeatureRepo devicefeature.Repository,
) devicefeature.UseCase {
	return &UseCase{
		logger: logger,
		eb:     eventBus,

		deviceRepo:        deviceRepo,
		deviceFeatureRepo: deviceFeatureRepo,
	}
}
