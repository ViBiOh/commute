package main

import (
	"net/http"

	"github.com/ViBiOh/commute/pkg/templ"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
)

func newPort(config configuration, clients clients, services services) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/token/{provider...}", services.commute.HandleToken())
	mux.Handle("/compute", services.commute.HandleCompute())
	mux.Handle("/", templ.Handler(*config.publicURL, services.strava.LoginURL()))

	return httputils.Handler(mux, clients.health,
		clients.telemetry.Middleware("http"),
		services.owasp.Middleware,
		services.cors.Middleware,
	)
}
