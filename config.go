package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetRoutes returns all the routes and their configs.
func GetRoutes() []ParentRoute {
	return []ParentRoute{
		{
			Path: "/users",
			SubRoutes: []Route{
				{
					path:           "/view/{id}",
					param:          "id",
					method:         http.MethodPost,
					contentType:    contentTypeJSON,
					responseStatus: http.StatusOK,
					response: marshallJSON(JSONResp{
						StatusCode: http.StatusOK,
						Message:    "Success",
					}),
				},
				{
					path:           "/edit/{id}",
					param:          "id",
					method:         http.MethodGet,
					contentType:    contentTypeJSON,
					responseStatus: http.StatusOK,
					response: marshallJSON(JSONResp{
						StatusCode: http.StatusOK,
						Message:    "User is authorized",
					}),
				},
			},
		},
	}
}

func marshallJSON(j JSONResp) []byte {
	payload, err := json.Marshal(j)
	if err != nil {
		_ = fmt.Errorf("failed creating route %v\n", j)
	}

	return payload
}
