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

func TestParams(t *testing.T) {
	var router = httprouter.New()
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var text = "Product " + params.ByName("id")
		fmt.Fprint(writer, text)
	})

	var request = httptest.NewRequest("GET", "http://localhost:8080/product/1", nil)
	var recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response = recorder.Result()
	var body, _ = io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
