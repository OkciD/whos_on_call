package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html
var templatesFS embed.FS

var CallStatus = template.Must(template.ParseFS(templatesFS, "base.html", "call_status.html"))
