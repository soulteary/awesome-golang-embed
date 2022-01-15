package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticRoute(t *testing.T) {
	router := registerRoute()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/assets/example.txt", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "@soulteary: Hello World", w.Body.String())
}

func TestSmallFileRepeatRequest(t *testing.T) {
	router := registerRoute()

	passed := true
	for i := 0; i < 1000; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/assets/vue.min.js", nil)
		router.ServeHTTP(w, req)

		if w.Code != 200 {
			passed = false
		}
	}

	assert.Equal(t, true, passed)
}

func TestLargeFileRepeatRequest(t *testing.T) {
	router := registerRoute()

	passed := true
	for i := 0; i < 100; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/assets/chip.jpg", nil)
		router.ServeHTTP(w, req)

		if w.Code != 200 {
			passed = false
		}
	}

	assert.Equal(t, true, passed)
}