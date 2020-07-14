package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const port = ":8000"
const (
	contentTypeJSON      = "application/json"
	contentTypePlainText = "text/plain"
)

func main() {
	rtr := mux.NewRouter()

	routes := GetRoutes()
	for _, sr := range routes {
		srtr := rtr.PathPrefix(sr.Path).Subrouter()

		for _, r := range sr.SubRoutes {
			srtr.HandleFunc(r.path, r.Handler).Methods(r.method)
		}
	}

	http.Handle("/", rtr)

	fmt.Println("Listening on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
