package sse

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"

	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/server/web/assets"
	"github.com/OkciD/whos_on_call/internal/shared/eventbus"
	"github.com/OkciD/whos_on_call/internal/shared/models"

	"go.jetify.com/sse"
)

type constants struct {
	CallStateInactive models.CallState
	CallStateActive   models.CallState
}

type callStatusTemplateData struct {
	CallStatus models.CallStatus

	Constants constants
}

var callStatusPartialTemplate = template.Must(template.ParseFS(assets.HTMLFiles, "partials/call_status.html"))

func (h *Handler) renderCallStatusPartial(ctx context.Context) ([]byte, error) {
	callStatus, err := h.callStatusUseCase.Calculate(ctx)
	if err != nil {
		return nil, fmt.Errorf("error calculating call status: %w", err)
	}

	templateData := callStatusTemplateData{
		CallStatus: callStatus,

		Constants: constants{
			CallStateInactive: models.CallStateInactive,
			CallStateActive:   models.CallStateActive,
		},
	}

	var buf bytes.Buffer

	if err := callStatusPartialTemplate.ExecuteTemplate(&buf, "partial:callstatus", templateData); err != nil {
		return nil, fmt.Errorf("error executing template for call status: %w", err)
	}

	return buf.Bytes(), nil
}

func (h *Handler) sendCallStatusToSSE(ctx context.Context, conn *sse.Conn) error {
	content, err := h.renderCallStatusPartial(ctx)
	if err != nil {
		return fmt.Errorf("failed to render call status partial: %w", err)
	}

	sseEvent := sse.Event{
		ID:    uuid.NewString(),
		Event: "CallStatus",
		Data:  sse.Raw(content),
		Split: true,
	}

	if err := conn.SendEvent(ctx, &sseEvent); err != nil {
		return fmt.Errorf("failed to send event: %w", err)
	}

	return nil
}

func (h *Handler) callStatus() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		connID := appContext.GetRequestID(r.Context())
		if connID == "" {
			connID = uuid.NewString()
		}

		logger := h.logger.WithField("conn_id", connID)

		conn, err := sse.Upgrade(r.Context(), w)
		if err != nil {
			logger.WithError(err).Error("failed to upgrade connection to sse")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		h.sseConnectionsMux.Lock()
		h.sseConnections[connID] = conn
		h.sseConnectionsMux.Unlock()

		defer func() {
			err := conn.Close()
			if err != nil {
				logger.WithError(err).Warn("failed to close sse connection")
			} else {
				logger.Debug("sse connection closed")
			}
			h.sseConnectionsMux.Lock()
			delete(h.sseConnections, connID)
			h.sseConnectionsMux.Unlock()
		}()

		logger.Info("sse conn established")

		callStatusUpdatedChan := h.eb.On(eventbus.EventTypeDeviceFeatureUpdated{})
		defer h.eb.Off(eventbus.EventTypeDeviceFeatureUpdated{}, callStatusUpdatedChan)

		if err := h.sendCallStatusToSSE(r.Context(), conn); err != nil {
			logger.WithError(err).Error("failed to send initial event")
			return
		}

		for {
			select {
			case <-callStatusUpdatedChan:
				logger.Debug("call status update event received from eventbus")
				if err := h.sendCallStatusToSSE(r.Context(), conn); err != nil {
					logger.WithError(err).Error("failed to send subsequent sse event")
					return
				}
				continue
			case <-r.Context().Done():
				return
			}
		}
	})
}
