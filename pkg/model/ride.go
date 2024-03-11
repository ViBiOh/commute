package model

import (
	"time"

	"github.com/ViBiOh/commute/pkg/coordinates"
)

const (
	HOME_ARRIVE = 1 << iota
	WORK_LEAVE
	WORK_ARRIVE
	HOME_LEAVE
	COMMUTE
)

type Ride struct {
	Date     time.Time
	Duration time.Duration
	Start    coordinates.LatLng
	End      coordinates.LatLng
	Commute  bool
}

type Rides []Ride

func (r Rides) Coordinates() coordinates.List {
	coords := make(coordinates.List, len(r)*2)

	for index, ride := range r {
		coords[index*2] = ride.Start
		coords[index*2+1] = ride.End
	}

	return coords
}

func (r Rides) ToCommutes(home, work coordinates.LatLng, distance float64) (Commutes, error) {
	roundTrips := Commutes{}

	for _, ride := range r {
		day := ride.Date.Format(time.DateOnly)

		if ride.Start.IsWithin(home, distance) {
			roundTrips[day] |= HOME_LEAVE
		}

		if ride.End.IsWithin(work, distance) {
			roundTrips[day] |= WORK_ARRIVE
		}

		if ride.Start.IsWithin(work, distance) {
			roundTrips[day] |= WORK_LEAVE
		}

		if ride.End.IsWithin(home, distance) {
			roundTrips[day] |= HOME_ARRIVE
		}

		if ride.Commute {
			roundTrips[day] |= COMMUTE
		}
	}

	return roundTrips, nil
}
