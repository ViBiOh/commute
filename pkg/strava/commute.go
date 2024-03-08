package strava

import (
	"context"
	"fmt"
	"log/slog"
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

type Rides []Ride

func (r Rides) Coordinates() coordinates.List {
	coords := make(coordinates.List, len(r)*2)

	for index, ride := range r {
		coords[index*2] = ride.start
		coords[index*2+1] = ride.end
	}

	return coords
}

func getRides(activities []Activity) (Rides, error) {
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
			slog.LogAttrs(context.Background(), slog.LevelError, "unable to parse start", slog.Any("error", err), slog.Any("activity", activity))
			return nil, fmt.Errorf("parse start: %w", err)
		}

		end, err := coordinates.NewLatLng(activity.EndLatlng)
		if err != nil {
			slog.LogAttrs(context.Background(), slog.LevelError, "unable to parse end", slog.Any("error", err), slog.Any("activity", activity))
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

func getCommutes(rides Rides, home, work coordinates.LatLng, distance float64) (model.Commutes, error) {
	roundTrips := model.Commutes{}

	for _, ride := range rides {
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
