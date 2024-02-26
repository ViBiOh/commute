package strava

import (
	"fmt"
	"time"

	"github.com/ViBiOh/strava/pkg/coordinates"
	"github.com/ViBiOh/strava/pkg/model"
	"github.com/twpayne/go-polyline"
)

const (
	HOME_ARRIVE = 1 << iota
	WORK_LEAVE
	WORK_ARRIVE
	HOME_LEAVE
	COMMUTE
)

type Ride struct {
	date    time.Time
	start   coordinates.LatLng
	end     coordinates.LatLng
	commute bool
}

func getClusters(rides []Ride) []coordinates.LatLng {
	var coords []coordinates.LatLng
	for _, ride := range rides {
		coords = append(coords, ride.start)
		coords = append(coords, ride.end)
	}

	var groups []coordinates.LatLng

	for len(coords) > 0 {
		var nextCoords []coordinates.LatLng

		current := coords[0]
		currentGroup := []coordinates.LatLng{current}

		for i := 1; i < len(coords); i++ {
			following := coords[i]

			if current.IsWithin(following, 0.5) {
				currentGroup = append(currentGroup, following)
			} else {
				nextCoords = append(nextCoords, following)
			}
		}

		groups = append(groups, coordinates.Center(currentGroup...))
		coords = nextCoords
	}

	return groups
}

func getRides(activities []Activity) ([]Ride, error) {
	var output []Ride

	for _, activity := range activities {
		if activity.Type != "Ride" {
			continue
		}

		if len(activity.Map.SummaryPolyline) != 0 {
			coords, _, err := polyline.DecodeCoords([]byte(activity.Map.SummaryPolyline))
			if err != nil {
				return nil, fmt.Errorf("decode polyline: %w", err)
			}

			if len(coords) > 2 {
				activity.StartLatlng = coords[0]
				activity.EndLatlng = coords[len(coords)-1]
			}
		}

		start, err := coordinates.NewLatLng(activity.StartLatlng)
		if err != nil {
			return nil, fmt.Errorf("parse start: %w", err)
		}

		end, err := coordinates.NewLatLng(activity.EndLatlng)
		if err != nil {
			return nil, fmt.Errorf("parse end: %w", err)
		}

		output = append(output, Ride{
			date:    activity.StartDate,
			start:   start,
			end:     end,
			commute: activity.Commute,
		})
	}

	return output, nil
}

func getCommutes(rides []Ride, home, work coordinates.LatLng, distance float64) (model.Commutes, error) {
	roundTrips := model.Commutes{}

	for _, ride := range rides {
		if weekday := ride.date.Weekday(); weekday < 0 || weekday > 5 {
			continue
		}

		day := ride.date.Format(time.DateOnly)

		if ride.start.IsWithin(home, distance) {
			roundTrips[day] |= HOME_LEAVE
		}

		if ride.end.IsWithin(work, distance) {
			roundTrips[day] |= WORK_ARRIVE
		}

		if ride.start.IsWithin(work, distance) {
			roundTrips[day] |= WORK_LEAVE
		}

		if ride.end.IsWithin(home, distance) {
			roundTrips[day] |= HOME_ARRIVE
		}

		if ride.commute {
			roundTrips[day] |= COMMUTE
		}
	}

	return roundTrips, nil
}
