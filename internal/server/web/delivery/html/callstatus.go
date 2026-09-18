package html

import (
	"html/template"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/server/web/assets"
	"github.com/OkciD/whos_on_call/internal/shared/models"
)

type constants struct {
	CallStateInactive models.CallState
	CallStateActive   models.CallState
}

type callStatusTemplateData struct {
	CallStatus models.CallStatus

	Constants constants
}

var callStatusPageTemplate = template.Must(template.ParseFS(
	assets.HTMLFiles,
	"base.html",
	"pages/call_status.html",
	"partials/call_status.html",
))
var callStatusPartialTemplate = template.Must(template.ParseFS(assets.HTMLFiles, "partials/call_status.html"))

func (h *WebHandler) callStatusPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if err := callStatusPageTemplate.ExecuteTemplate(w, "base", nil); err != nil {
			h.logger.WithError(err).Error("error executing template for call status")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func (h *WebHandler) callStatusPartial() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callStatus, err := h.callStatusUseCase.Calculate(r.Context())
		if err != nil {
			h.logger.WithError(err).Error("error calculating call status")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		templateData := callStatusTemplateData{
			CallStatus: callStatus,

			Constants: constants{
				CallStateInactive: models.CallStateInactive,
				CallStateActive:   models.CallStateActive,
			},
		}

		if err := callStatusPartialTemplate.ExecuteTemplate(w, "partial:callstatus", templateData); err != nil {
			h.logger.WithError(err).Error("error executing template for call status")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
