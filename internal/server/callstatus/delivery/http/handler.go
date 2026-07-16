package http

import (
	"github.com/OkciD/whos_on_call/internal/server/callstatus"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type CallStatusHandler struct {
	logger logger.Logger

	callStatusUseCase callstatus.UseCase
}

func New(logger logger.Logger, callStatusUseCase callstatus.UseCase) CallStatusHandler {
	return CallStatusHandler{
		logger: logger,

		callStatusUseCase: callStatusUseCase,
	}
}
