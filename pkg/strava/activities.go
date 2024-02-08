package strava

import (
	"context"
	"fmt"

	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

func (s Service) getActivities(ctx context.Context, requester request.Request) ([]Activity, error) {
	resp, err := requester.Path("/athlete/activities?per_page=100").Send(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("get activities: %w", err)
	}

	var activities []Activity
	if err = httpjson.Read(resp, &activities); err != nil {
		return nil, err
	}

	return activities, nil
}
