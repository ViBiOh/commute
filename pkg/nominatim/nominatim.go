package nominatim

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/strava/pkg/coordinates"
)

var (
	requester   = request.New().URL("https://nominatim.openstreetmap.org")
	ErrNotFound = errors.New("no address found")
)

type SearchResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GetLatLng(ctx context.Context, query string) (coordinates.LatLng, error) {
	urlQuery := url.Values{}
	urlQuery.Add("format", "json")
	urlQuery.Add("q", query)

	resp, err := requester.Path("/search?"+urlQuery.Encode()).Send(ctx, nil)
	if err != nil {
		return coordinates.LatLng{}, fmt.Errorf("search: %w", err)
	}

	var responses []SearchResponse
	if err := httpjson.Read(resp, &responses); err != nil {
		return coordinates.LatLng{}, fmt.Errorf("parse: %w", err)
	}

	if len(responses) == 0 {
		return coordinates.LatLng{}, ErrNotFound
	}

	lat, err := strconv.ParseFloat(responses[0].Lat, 64)
	if err != nil {
		return coordinates.LatLng{}, fmt.Errorf("parse lat `%s`: %w", responses[0].Lat, err)
	}

	lng, err := strconv.ParseFloat(responses[0].Lon, 64)
	if err != nil {
		return coordinates.LatLng{}, fmt.Errorf("parse lon `%s`: %w", responses[0].Lon, err)
	}

	return coordinates.LatLng{lat, lng}, nil
}
