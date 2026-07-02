package http

import (
	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type Handler struct {
	logger logger.Logger

	deviceFeatureUseCase devicefeature.UseCase
}

func New(logger logger.Logger, deviceFeatureUseCase devicefeature.UseCase) *Handler {
	h := &Handler{
		logger: logger,

		deviceFeatureUseCase: deviceFeatureUseCase,
	}

	return h
}
