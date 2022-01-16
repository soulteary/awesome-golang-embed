package main

//go:generate esc  -o assets/assets.go -pkg assets -prefix "assets/"  assets

import (
	"log"
	"net/http"
	"solution-embed/assets"
)

func main() {
	mutex := http.NewServeMux()
	// FS() is created by esc and returns a http.Filesystem.
	mutex.Handle("/", http.FileServer(assets.FS(false)))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
