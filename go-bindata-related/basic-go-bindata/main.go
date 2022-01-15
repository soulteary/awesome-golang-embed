package main

import (
	"log"
	"net/http"

	"solution-embed/assets"
)

//go:generate go-bindata -fs -o=assets/assets.go -pkg=assets ./assets

func main() {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(assets.AssetFile()))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}