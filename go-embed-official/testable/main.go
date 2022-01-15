package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed assets
var assets embed.FS

func registerRoute() *http.ServeMux {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(assets)))
	return mutex
}

func main() {
	mutex := registerRoute()
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
