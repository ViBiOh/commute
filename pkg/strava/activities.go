package strava

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/strava/pkg/coordinates"
	"github.com/ViBiOh/strava/pkg/model"
	"github.com/twpayne/go-polyline"
)

func (s Service) getActivities(ctx context.Context, requester request.Request, before, after time.Time) ([]Activity, error) {
	params := url.Values{}
	params.Add("per_page", "100")

	if !before.IsZero() {
		params.Add("before", strconv.FormatInt(before.Unix(), 10))
	}

	if !after.IsZero() {
		params.Add("after", strconv.FormatInt(after.Unix(), 10))
	}

	resp, err := requester.Path("/athlete/activities?%s", params.Encode()).Send(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("get activities: %w", err)
	}

	var activities []Activity
	if err = httpjson.Read(resp, &activities); err != nil {
		return nil, err
	}

	return activities, nil
}

func toRides(activities []Activity) (model.Rides, error) {
	var output []model.Ride

	for _, activity := range activities {
		if activity.Type != "Ride" {
			continue
		}

		if len(activity.Map.SummaryPolyline) != 0 {
			coords, _, err := polyline.DecodeCoords([]byte(activity.Map.SummaryPolyline))
			if err != nil {
				return nil, fmt.Errorf("decode polyline: %w", err)
			}

			if len(coords) > 2 && len(coords[0]) > 0 && len(coords[len(coords)-1]) > 0 {
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

		output = append(output, model.Ride{
			Date:    activity.StartDate,
			Start:   start,
			End:     end,
			Commute: activity.Commute,
		})
	}

	return output, nil
}
