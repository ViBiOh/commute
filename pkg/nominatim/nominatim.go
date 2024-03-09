package nominatim

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ViBiOh/commute/pkg/coordinates"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

var (
	requester   = request.New().URL("https://nominatim.openstreetmap.org")
	ErrNotFound = errors.New("no address found")
)

type Response struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

func Geocode(ctx context.Context, query string) (coordinates.LatLng, string, error) {
	urlQuery := url.Values{}
	urlQuery.Add("format", "json")
	urlQuery.Add("q", query)

	resp, err := requester.Path("/search?"+urlQuery.Encode()).Send(ctx, nil)
	if err != nil {
		return coordinates.LatLng{}, "", fmt.Errorf("search: %w", err)
	}

	var responses []Response
	if err := httpjson.Read(resp, &responses); err != nil {
		return coordinates.LatLng{}, "", fmt.Errorf("parse: %w", err)
	}

	if len(responses) == 0 {
		return coordinates.LatLng{}, "", ErrNotFound
	}

	lat, err := strconv.ParseFloat(responses[0].Lat, 64)
	if err != nil {
		return coordinates.LatLng{}, "", fmt.Errorf("parse lat `%s`: %w", responses[0].Lat, err)
	}

	lng, err := strconv.ParseFloat(responses[0].Lon, 64)
	if err != nil {
		return coordinates.LatLng{}, "", fmt.Errorf("parse lon `%s`: %w", responses[0].Lon, err)
	}

	return coordinates.LatLng{lat, lng}, responses[0].DisplayName, nil
}

func Reverse(ctx context.Context, coord coordinates.LatLng) (string, error) {
	urlQuery := url.Values{}
	urlQuery.Add("format", "json")
	urlQuery.Add("zoom", "16") // borough
	urlQuery.Add("lat", fmt.Sprintf("%f", coord[0]))
	urlQuery.Add("lon", fmt.Sprintf("%f", coord[1]))

	resp, err := requester.Path("/reverse?"+urlQuery.Encode()).Send(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("reverse: %w", err)
	}

	var response Response
	if err := httpjson.Read(resp, &response); err != nil {
		return "", fmt.Errorf("parse: %w", err)
	}

	return response.DisplayName, nil
}
