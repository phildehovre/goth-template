//go:build !dev
// +build !dev

package main

import (
	"embed"
	"fmt"
	"net/http"
)

var publicFS embed.FS

func public() http.Handler {
	fmt.Println("From the static_prod file")
	return http.FileServerFS(publicFS)
}
