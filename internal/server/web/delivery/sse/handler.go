package sse

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/server/callstatus"
	"github.com/OkciD/whos_on_call/internal/server/pkg/sse"
	"github.com/OkciD/whos_on_call/internal/shared/eventbus"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type Handler struct {
	logger logger.Logger

	eb eventbus.EventBus

	ssePool sse.Pool

	callStatusUseCase callstatus.UseCase
}

func New(
	mux *http.ServeMux,
	logger logger.Logger,
	eventBus eventbus.EventBus,
	callStatusUseCase callstatus.UseCase,
) *Handler {
	h := &Handler{
		logger: logger,

		eb: eventBus,

		ssePool: sse.NewPool(logger),

		callStatusUseCase: callStatusUseCase,
	}

	mux.Handle("GET /sse/callstatus", h.callStatus())

	return h
}

func (h *Handler) Stop() {
	h.ssePool.CloseAllConns()
}
