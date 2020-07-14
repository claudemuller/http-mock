package main

import "net/http"

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
					response: jsonResp{
						StatusCode: http.StatusOK,
						Message:    "Success",
					},
				},
				{
					path:           "/edit/{id}",
					param:          "id",
					method:         http.MethodGet,
					contentType:    contentTypeJSON,
					responseStatus: http.StatusOK,
					response: jsonResp{
						StatusCode: http.StatusOK,
						Message:    "User is authorized",
					},
				},
			},
		},
	}
}
