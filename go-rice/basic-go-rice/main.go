package main

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func main() {
	conf := rice.Config{
		LocateOrder: []rice.LocateMethod{rice.LocateEmbedded, rice.LocateAppended, rice.LocateFS},
	}
	box, err := conf.FindBox("assets")
	if err != nil {
		log.Fatalf("error opening rice.Box: %s\n", err)
	}

	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(box.HTTPBox()))
	err = http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
