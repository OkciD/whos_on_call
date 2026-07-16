package http

import (
	"github.com/OkciD/whos_on_call/internal/server/user"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type UserHandler struct {
	logger logger.Logger

	userUseCase user.UseCase
}

func New(logger logger.Logger, userUseCase user.UseCase) UserHandler {
	return UserHandler{
		logger: logger,

		userUseCase: userUseCase,
	}
}
