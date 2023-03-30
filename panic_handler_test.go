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

func TestPanicHandler(t *testing.T) {
	var router = httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic : ", error)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Ups")
	})

	var request = httptest.NewRequest("GET", "http://localhost:8080/", nil)
	var recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response = recorder.Result()
	var body, _ = io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(body))
}
