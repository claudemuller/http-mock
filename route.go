package main

import (
	"encoding/json"
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
	response       jsonResp
}

type jsonResp struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

// Handler handles the request and sends responses.
func (route *Route) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params[route.param]
	payload := []byte("")

	var err error

	fmt.Printf("Received a %s from %s to %s\n", r.Method, r.RemoteAddr, r.RequestURI)
	fmt.Printf("Param %s:, %s\n", route.param, param)

	var contentType string

	if route.contentType == contentTypeJSON {
		payload, err = json.Marshal(route.response)
		if err != nil {
			errMsg := "There was a problem crafting the json"
			fmt.Println(errMsg)
			http.Error(w, errMsg, http.StatusInternalServerError)

			return
		}

		contentType = route.contentType
	} else {
		contentType = contentTypePlainText
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(route.responseStatus)

	written, err := w.Write(payload)
	if err != nil {
		fmt.Printf("Error writing response: %s\n", err)
	}

	fmt.Printf("%d bytes written\n\n", written)
}
