package main

import (
	"github.com/ViBiOh/strava/pkg/mapbox"
	"github.com/ViBiOh/strava/pkg/strava"
)

type service struct {
	mapbox mapbox.Service
	strava strava.Service
}

func newService(config configuration) service {
	mapboxService := mapbox.New(config.mapbox)

	return service{
		mapbox: mapboxService,
		strava: strava.New(config.strava, *config.publicURL, mapboxService),
	}
}
