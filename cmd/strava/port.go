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
	mux.Handle("/", templ.Handler(*config.publicURL, service.strava.StravaLoginURL()))

	return mux
}
