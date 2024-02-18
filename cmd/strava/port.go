package main

import (
	"net/http"

	"github.com/ViBiOh/strava/pkg/templ"
)

const (
	apiPath = "/api/{action}"
)

func newPort(config configuration, service service) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(apiPath, service.strava.Handle())
	mux.Handle("/", http.HandlerFunc(templ.HandlerFunc))

	return mux
}
