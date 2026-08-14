package quizforge

import (
	"embed"
	"io/fs"
)

//go:embed all:static
var staticFS embed.FS

func StaticFS() fs.FS {
	f, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
	return f
}
