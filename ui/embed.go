package ui

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:generate npm run build
//go:embed dist/*
var files embed.FS

func SvelteKitFS() http.FileSystem {
	fsys, err := fs.Sub(files, "dist")
	if err != nil {
		log.Fatal(err)
	}
	filesystem := http.FS(fsys)
	return filesystem
}
