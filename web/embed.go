package web

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed build/*
var FS embed.FS

var BuildFS fs.FS

func init() {
	buildFS, err := fs.Sub(FS, "build")
	if err != nil {
		log.Fatalln(err)
	}
	BuildFS = buildFS
}
