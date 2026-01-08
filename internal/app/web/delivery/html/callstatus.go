package html

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/app/web/templates"
)

type constants struct {
	CallStateInactive models.CallState
	CallStateActive   models.CallState
}

type callStatusTemplateData struct {
	CallStatus models.CallStatus

	Constants constants
}

func (h *WebHandler) CallStatus() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callStatus, err := h.callStatusUseCase.Calculate(r.Context())
		if err != nil {
			h.logger.WithError(err).Error("error calculating call status")
			w.WriteHeader(http.StatusInternalServerError)
		}

		templateData := callStatusTemplateData{
			CallStatus: callStatus,

			Constants: constants{
				CallStateInactive: models.CallStateInactive,
				CallStateActive:   models.CallStateActive,
			},
		}

		if err := templates.CallStatus.Execute(w, templateData); err != nil {
			h.logger.WithError(err).Error("error executing template for call status")
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
