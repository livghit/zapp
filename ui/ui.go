package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var distFiles embed.FS

func SpaHandler() (http.Handler, error) {
	fsys, err := fs.Sub(distFiles, "dist")
	if err != nil {
		return nil, err
	}

	return http.FileServer(http.FS(fsys)), nil
}
