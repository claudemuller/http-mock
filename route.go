package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ParentRoute struct {
	Path      string
	SubRoutes []Route
}

// Route defines the config for a route.
type Route struct {
	path           string
	param          string
	method         string
	contentType    string
	responseStatus int
	response       []byte
}

type JSONResp struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

// Handler handles the request and sends responses.
func (route *Route) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params[route.param]

	var err error

	fmt.Printf("Received a %s from %s to %s\n", r.Method, r.RemoteAddr, r.RequestURI)
	fmt.Printf("Param %s:, %s\n", route.param, param)

	setContentType(w, route.contentType)
	w.WriteHeader(route.responseStatus)

	written, err := w.Write(route.response)
	if err != nil {
		fmt.Printf("Error writing response: %s\n", err)
	}

	fmt.Printf("%d bytes written\n\n", written)
}

func setContentType(w http.ResponseWriter, ct string) {
	var contentType string

	switch ct {
	case contentTypeJSON:
		contentType = ct
	default:
		contentType = contentTypePlainText
	}

	w.Header().Set("Content-Type", contentType)
}
