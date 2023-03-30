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

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(writer, request)

}

func TestMiddleware(t *testing.T) {
	var router = httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Middleware")
	})

	var middleware = LogMiddleware{Handler: router}

	var request = httptest.NewRequest("GET", "http://localhost:8080/", nil)
	var recorder = httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	var response = recorder.Result()
	var body, _ = io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
}
