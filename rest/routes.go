package rest

import (
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

// Endpoint is a api endpoint for serving a rest api, ex: get /api/v1/cars
type Endpoint struct {
	Version  int
	Path     string
	Method   string
	Callback HandleFunc
}

// Route is a collection of endpoints
type Route struct {
	Name        string
	Entrypoints []Endpoint
	InitFunc    func()
}
