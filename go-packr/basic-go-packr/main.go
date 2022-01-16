package main

import (
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("myBox", "./assets")

	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(box))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
