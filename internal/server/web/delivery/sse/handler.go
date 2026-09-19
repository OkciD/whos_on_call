package sse

import (
	"net/http"
	"sync"

	"github.com/OkciD/whos_on_call/internal/server/callstatus"
	"github.com/OkciD/whos_on_call/internal/shared/eventbus"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"

	"go.jetify.com/sse"
)

type Handler struct {
	logger logger.Logger

	eb eventbus.EventBus

	sseConnectionsMux sync.Mutex
	sseConnections    map[string]*sse.Conn

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

		sseConnections: make(map[string]*sse.Conn),

		callStatusUseCase: callStatusUseCase,
	}

	mux.Handle("GET /sse/callstatus", h.callStatus())

	return h
}

func (h *Handler) Stop() {
	h.sseConnectionsMux.Lock()
	defer h.sseConnectionsMux.Unlock()

	for id, conn := range h.sseConnections {
		logger := h.logger.WithField("conn_id", id)

		err := conn.Close()
		if err != nil {
			logger.WithError(err).Warn("failed to close sse connection")
		}
		logger.Debug("sse connection closed")
	}

	h.sseConnections = make(map[string]*sse.Conn, 0)
}
