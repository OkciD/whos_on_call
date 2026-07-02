package http

import (
	"github.com/OkciD/whos_on_call/internal/server/user"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type Handler struct {
	logger logger.Logger

	userUseCase user.UseCase
}

func New(logger logger.Logger, userUseCase user.UseCase) *Handler {
	h := &Handler{
		logger: logger,

		userUseCase: userUseCase,
	}

	return h
}
