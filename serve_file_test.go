package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFiles(t *testing.T) {
	var router = httprouter.New()
	var directory, _ = fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	var request = httptest.NewRequest("GET", "http://localhost:8080/files/goodbye.txt", nil)
	var recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response = recorder.Result()
	var body, _ = io.ReadAll(response.Body)

	assert.Equal(t, "Goodbye Httprouter", string(body))
}
