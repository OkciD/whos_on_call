package usecase

import (
	"github.com/OkciD/whos_on_call/internal/pkg/duration"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	"github.com/OkciD/whos_on_call/internal/server/callstatus"
	"github.com/OkciD/whos_on_call/internal/server/device"
	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
	"github.com/OkciD/whos_on_call/internal/server/user"
)

type Config struct {
	RelaxationPeriod duration.MarshallableDuration `json:"relaxationPeriod"`
}

type UseCase struct {
	logger logger.Logger

	config Config

	userRepo          user.Repository
	deviceRepo        device.Repository
	deviceFeatureRepo devicefeature.Repository
}

func New(
	logger logger.Logger,
	config Config,
	userRepo user.Repository,
	deviceRepo device.Repository,
	deviceFeatureRepo devicefeature.Repository,
) callstatus.UseCase {
	return &UseCase{
		logger: logger,

		config: config,

		userRepo:          userRepo,
		deviceRepo:        deviceRepo,
		deviceFeatureRepo: deviceFeatureRepo,
	}
}
