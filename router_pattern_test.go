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

func TestRouterPatternNamedParameter(t *testing.T) {
	var router = httprouter.New()
	router.GET("/product/:id/item/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var text = "Product " + params.ByName("id") + " Item " + params.ByName("itemId")
		fmt.Fprint(writer, text)
	})

	var request = httptest.NewRequest("GET", "http://localhost:8080/product/1/item/1", nil)
	var recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response = recorder.Result()
	var body, _ = io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	var router = httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var text = "Image : " + params.ByName("image")
		fmt.Fprint(writer, text)
	})

	var request = httptest.NewRequest("GET", "http://localhost:8080/images/small/profile.png", nil)
	var recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response = recorder.Result()
	var body, _ = io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}