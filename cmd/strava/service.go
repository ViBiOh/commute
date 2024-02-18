package main

import "github.com/ViBiOh/strava/pkg/strava"

type service struct {
	strava strava.Service
}

func newService(config configuration) service {
	return service{
		strava: strava.New(config.strava, *config.publicURL),
	}
}
