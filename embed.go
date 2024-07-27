package frontend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

// Embed the build directory from the frontend.
//
//go:embed ui/dist/*
//go:embed ui/dist/assets/*
var BuildFs embed.FS

// Get the subtree of the embedded files with `build` directory as a root.
func UI() http.FileSystem {
	build, err := fs.Sub(BuildFs, "ui/dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(build)
}
