package handlers

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed tmpls/*
var FS embed.FS

var TmplFS fs.FS

func init() {
	tmplFS, err := fs.Sub(FS, "tmpls")
	if err != nil {
		log.Fatalln(err)
	}
	TmplFS = tmplFS
}
