package main

import (
	"net/http"

	"github.com/ViBiOh/commute/pkg/templ"
)

func newPort(config configuration, service service) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/token/{provider...}", service.commute.HandleToken())
	mux.Handle("/compute", service.commute.HandleCompute())
	mux.Handle("/", templ.Handler(*config.publicURL, service.strava.LoginURL()))

	return mux
}
