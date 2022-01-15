package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"runtime"

	"solution-embed/pkg/assets"
)

//go:generate go-bindata -fs -o=pkg/assets/assets.go -pkg=assets ./assets

func registerRoute() *http.ServeMux {

	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(assets.AssetFile()))
	return mutex
}

func enableProf(mutex *http.ServeMux) {
	runtime.GOMAXPROCS(2)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	mutex.HandleFunc("/debug/pprof/", pprof.Index)
	mutex.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mutex.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mutex.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mutex.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func main() {
	mutex := registerRoute()
	enableProf(mutex)

	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
