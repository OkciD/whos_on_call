package http

import (
	"github.com/OkciD/whos_on_call/internal/server/callstatus"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type Handler struct {
	logger logger.Logger

	callStatusUseCase callstatus.UseCase
}

func New(logger logger.Logger, callStatusUseCase callstatus.UseCase) *Handler {
	h := &Handler{
		logger: logger,

		callStatusUseCase: callStatusUseCase,
	}

	return h
}
