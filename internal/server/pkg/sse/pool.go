package sse

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"go.jetify.com/sse"

	"github.com/google/uuid"

	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type Pool interface {
	NewConn(w http.ResponseWriter, r *http.Request) (*Conn, error)
	NewConnWithID(id string, w http.ResponseWriter, r *http.Request) (*Conn, error)
	CloseConn(id string)
	CloseAllConns()
}

type pool struct {
	logger logger.Logger

	mu    sync.Mutex
	conns map[string]*Conn

	closed bool
}

func NewPool(logger logger.Logger) Pool {
	return &pool{
		conns:  make(map[string]*Conn),
		logger: logger,
	}
}

func (h *pool) NewConn(w http.ResponseWriter, r *http.Request) (*Conn, error) {
	id := uuid.NewString()
	conn, err := h.NewConnWithID(id, w, r)
	return conn, err
}

func (h *pool) NewConnWithID(
	id string,
	w http.ResponseWriter,
	r *http.Request,
) (*Conn, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.conns[id]; ok {
		return nil, fmt.Errorf("sse conn id %s already exists", id)
	}

	conn, err := sse.Upgrade(r.Context(), w)
	if err != nil {
		return nil, fmt.Errorf("failed to upgrade %s connection to sse: %w", id, err)
	}

	ctx, cancel := context.WithCancel(r.Context())

	h.conns[id] = &Conn{
		Conn: conn,

		ID:       id,
		ctx:      ctx,
		cancelFn: cancel,
	}

	return h.conns[id], nil
}

func (h *pool) CloseConn(id string) {
	logger := h.logger.WithField("conn_id", id)

	if h.closed {
		logger.Debug("connection already closed by CloseAllConns")
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	conn, ok := h.conns[id]
	if !ok {
		logger.Warn("tried closing conn by non-existing id")
		return
	}

	err := conn.Close()
	if err != nil {
		logger.WithError(err).Warn("failed to close sse connection")
	}

	conn.cancelFn()

	logger.Debug("sse connection closed")

	delete(h.conns, id)
}

func (h *pool) CloseAllConns() {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.closed = true

	for id, conn := range h.conns {
		logger := h.logger.WithField("conn_id", id)

		err := conn.Close()
		if err != nil {
			logger.WithError(err).Warn("failed to close sse connection")
		}

		conn.cancelFn()

		logger.Debug("sse connection closed")
	}

	h.conns = make(map[string]*Conn)
}
