package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	"github.com/OkciD/whos_on_call/internal/server/callstatus"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
)

type CallStatusHandler struct {
	logger logger.Logger

	callStatusUseCase callstatus.UseCase
}

func New(mux *http.ServeMux, logger logger.Logger, callStatusUseCase callstatus.UseCase) *CallStatusHandler {
	h := &CallStatusHandler{
		logger: logger,

		callStatusUseCase: callStatusUseCase,
	}

	mux.Handle("GET /api/v1/status", handler.GenericHandler(logger, h.GetStatus))

	return h
}
