package main

import (
	"github.com/ViBiOh/commute/pkg/commute"
	"github.com/ViBiOh/commute/pkg/mapbox"
	"github.com/ViBiOh/commute/pkg/strava"
	"github.com/ViBiOh/commute/pkg/wahoo"
)

type services struct {
	commute commute.Service
	strava  strava.Service
	wahoo   wahoo.Service
}

func newServices(config configuration) services {
	stravaService := strava.New(config.strava, *config.publicURL)
	wahooService := wahoo.New(config.wahoo, *config.publicURL)

	return services{
		strava:  stravaService,
		wahoo:   wahooService,
		commute: commute.New(*config.publicURL, mapbox.New(config.mapbox), stravaService, wahooService),
	}
}
