package strava

import (
	"fmt"
	"time"

	"github.com/ViBiOh/strava/pkg/coordinates"
	"github.com/ViBiOh/strava/pkg/model"
)

const (
	HOME_ARRIVE = 1 << iota
	WORK_LEAVE
	WORK_ARRIVE
	HOME_LEAVE
)

func computeCommute(activities []Activity, home, work coordinates.LatLng) (model.Commutes, error) {
	roundTrips := model.Commutes{}

	for _, activity := range activities {
		if weekday := activity.StartDate.Weekday(); activity.Type != "Ride" || weekday < 0 || weekday > 5 {
			continue
		}

		day := activity.StartDate.Format(time.DateOnly)

		var found bool

		startLatLng, err := coordinates.NewLatLng(activity.StartLatlng)
		if err != nil {
			return nil, fmt.Errorf("parse start: %w", err)
		}

		endLatLng, err := coordinates.NewLatLng(activity.EndLatlng)
		if err != nil {
			return nil, fmt.Errorf("parse end: %w", err)
		}

		if startLatLng.IsWithin(home, .5) {
			roundTrips[day] |= HOME_LEAVE
			found = true
		}

		if endLatLng.IsWithin(work, .5) {
			roundTrips[day] |= WORK_ARRIVE
			found = true
		}

		if startLatLng.IsWithin(work, .5) {
			roundTrips[day] |= WORK_LEAVE
			found = true
		}

		if endLatLng.IsWithin(home, .5) {
			roundTrips[day] |= HOME_ARRIVE
			found = true
		}

		if !found {
			fmt.Println(activity.StartDate, activity.Name, "from", activity.StartLatlng, "to", activity.EndLatlng)
		}
	}

	return roundTrips, nil
}
