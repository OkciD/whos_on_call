package htmx

import (
	"html/template"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/server/web/assets"
)

var callStatusPageTemplate = template.Must(template.ParseFS(
	assets.HTMLFiles,
	"base.html",
	"pages/call_status.html",
	"partials/call_status.html",
))

func (h *Handler) callStatusPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if err := callStatusPageTemplate.ExecuteTemplate(w, "base", nil); err != nil {
			h.logger.WithError(err).Error("error executing template for call status")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
