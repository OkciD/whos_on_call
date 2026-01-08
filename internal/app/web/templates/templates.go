package templates

import (
	"embed"
	"text/template"
)

//go:embed *.html
var templatesFS embed.FS

var CallStatus *template.Template = template.Must(template.ParseFS(templatesFS, "base.html", "call_status.html"))
