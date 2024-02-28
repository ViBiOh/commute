package strava

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
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
