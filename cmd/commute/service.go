package main

import (
	"github.com/ViBiOh/commute/pkg/commute"
	"github.com/ViBiOh/commute/pkg/mapbox"
	"github.com/ViBiOh/commute/pkg/strava"
)

type service struct {
	commute commute.Service
	strava  strava.Service
}

func newService(config configuration) service {
	stravaService := strava.New(config.strava, *config.publicURL)

	return service{
		strava:  stravaService,
		commute: commute.New(*config.publicURL, stravaService, mapbox.New(config.mapbox)),
	}
}
