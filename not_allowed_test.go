package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotAllowedHandler(t *testing.T) {
	var router = httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Gak boleh")
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "POST")
	})

	var request = httptest.NewRequest("GET", "http://localhost:8080/", nil)
	var recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response = recorder.Result()
	var body, _ = io.ReadAll(response.Body)

	assert.Equal(t, "Gak boleh", string(body))
}
