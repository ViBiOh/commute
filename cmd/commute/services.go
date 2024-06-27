package main

import (
	"github.com/ViBiOh/commute/pkg/commute"
	"github.com/ViBiOh/commute/pkg/mapbox"
	"github.com/ViBiOh/commute/pkg/strava"
	"github.com/ViBiOh/commute/pkg/wahoo"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

type services struct {
	server  *server.Server
	cors    cors.Service
	strava  strava.Service
	wahoo   wahoo.Service
	commute commute.Service
	owasp   owasp.Service
}

func newServices(config configuration) services {
	var output services

	output.server = server.New(config.server)
	output.owasp = owasp.New(config.owasp)
	output.cors = cors.New(config.cors)

	output.strava = strava.New(config.strava, *config.publicURL)
	output.wahoo = wahoo.New(config.wahoo, *config.publicURL)
	output.commute = commute.New(*config.publicURL, mapbox.New(config.mapbox), output.strava, output.wahoo)

	return output
}
