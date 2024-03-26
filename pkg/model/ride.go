package model

import (
	"sort"
	"time"

	"github.com/ViBiOh/commute/pkg/coordinates"
)

const (
	HomeArrive = 1 << iota
	WorkLeave
	WorkArrive
	HomeLeave
	Commute
)

type Ride struct {
	Date     time.Time
	ID       string
	Start    coordinates.LatLng
	End      coordinates.LatLng
	Duration time.Duration
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
	roundTrips := make(map[time.Time]*Day)

	for _, ride := range r {
		date := ride.Date.Truncate(time.Hour * 24)

		day, ok := roundTrips[date]
		if !ok {
			day = &Day{Date: date}
			roundTrips[date] = day
		}

		day.IDs = append(day.IDs, ride.ID)

		if ride.Start.IsWithin(home, distance) {
			day.Status |= HomeLeave
		}

		if ride.End.IsWithin(work, distance) {
			day.Status |= WorkArrive
		}

		if ride.Start.IsWithin(work, distance) {
			day.Status |= WorkLeave
		}

		if ride.End.IsWithin(home, distance) {
			day.Status |= HomeArrive
		}

		if ride.Commute {
			day.Status |= Commute
		}
	}

	return toCommutes(roundTrips), nil
}

func toCommutes(input map[time.Time]*Day) Commutes {
	output := make([]Day, 0, len(input))

	for _, value := range input {
		output = append(output, *value)
	}

	sort.Sort(CommutesByDate(output))

	return output
}
