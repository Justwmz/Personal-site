//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
	"os"
)

func public() http.Handler {
	fmt.Println("Serving a static files.")
	return http.StripPrefix("/static/", http.FileServerFS(os.DirFS("static")))
}
