package assets

import (
	"embed"
	"io/fs"

	"github.com/OkciD/whos_on_call/internal/shared/pkg/utils"
)

//go:embed "html" "static"
var files embed.FS

var (
	HTMLFiles   = sub(files, "html")
	StaticFiles = sub(files, "static")
)

func sub(f embed.FS, dir string) fs.FS {
	return utils.Must(fs.Sub(f, dir))
}
