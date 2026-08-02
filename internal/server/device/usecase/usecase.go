package usecase

import (
	"github.com/OkciD/whos_on_call/internal/server/device"
	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
	"github.com/OkciD/whos_on_call/internal/server/pkg/db"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type UseCase struct {
	logger logger.Logger

	txManager         db.TxManager
	deviceRepo        device.Repository
	deviceFeatureRepo devicefeature.Repository
}

func New(
	logger logger.Logger,
	txManager db.TxManager,
	deviceRepo device.Repository,
	deviceFeatureRepo devicefeature.Repository,
) device.UseCase {
	return &UseCase{
		logger: logger,

		txManager:         txManager,
		deviceRepo:        deviceRepo,
		deviceFeatureRepo: deviceFeatureRepo,
	}
}
