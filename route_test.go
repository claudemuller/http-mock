package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Status_Code(t *testing.T) {
	// Arrange
	route := Route{
		path:           "/ping",
		method:         http.MethodGet,
		contentType:    contentTypeJSON,
		responseStatus: http.StatusOK,
		response: marshallJSON(JSONResp{
			StatusCode: http.StatusOK,
			Message:    "Pong",
		}),
	}
	w, r := createHTTPReaderWriter(route.method, route.path, t)
	h := http.HandlerFunc(route.Handler)

	// Act
	h.ServeHTTP(w, r)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_ContentType_JSON(t *testing.T) {
	// Arrange
	route := Route{
		path:           "/ping",
		method:         http.MethodGet,
		contentType:    contentTypeJSON,
		responseStatus: http.StatusOK,
		response: marshallJSON(JSONResp{
			StatusCode: http.StatusOK,
			Message:    "Pong",
		}),
	}
	w, r := createHTTPReaderWriter(route.method, route.path, t)
	h := http.HandlerFunc(route.Handler)

	// Act
	h.ServeHTTP(w, r)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Reading from body failed: %v", err)
	}

	var res JSONResp

	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Fatalf("Unable to marshall json: %v", err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "Pong", res.Message)
}

func Test_ContentType_PlainText(t *testing.T) {
	// Arrange
	route := Route{
		path:           "/ping",
		method:         http.MethodGet,
		contentType:    contentTypePlainText,
		responseStatus: http.StatusOK,
		response:       []byte("Pong"),
	}
	w, r := createHTTPReaderWriter(route.method, route.path, t)
	h := http.HandlerFunc(route.Handler)

	// Act
	h.ServeHTTP(w, r)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Reading from body failed: %v", err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	assert.Equal(t, "Pong", string(body))
}

func Test_ContentType_Default(t *testing.T) {
	// Arrange
	route := Route{
		path:           "/ping",
		method:         http.MethodGet,
		contentType:    "something/weird",
		responseStatus: http.StatusOK,
		response:       []byte("Pong"),
	}
	w, r := createHTTPReaderWriter(route.method, route.path, t)
	h := http.HandlerFunc(route.Handler)

	// Act
	h.ServeHTTP(w, r)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Reading from body failed: %v", err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	assert.Equal(t, "Pong", string(body))
}

func createHTTPReaderWriter(method string, route string, t *testing.T) (*httptest.ResponseRecorder, *http.Request) {
	r, err := http.NewRequest(method, route, nil)
	if err != nil {
		t.Fatalf("Unable to create request: %v", err)
	}

	w := httptest.NewRecorder()

	return w, r
}
