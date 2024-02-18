package strava

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ViBiOh/strava/pkg/coordinates"
)

const keyName = "key"

func encodeKey(home, work coordinates.LatLng) string {
	redirectValues := url.Values{}
	redirectValues.Add(keyName, base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s|%s", home.String(), work.String()))))

	return redirectValues.Encode()
}

func parseKey(r *http.Request) (coordinates.LatLng, coordinates.LatLng, error) {
	rawKey, err := base64.StdEncoding.DecodeString(r.URL.Query().Get(keyName))
	if err != nil {
		return coordinates.LatLng{}, coordinates.LatLng{}, fmt.Errorf("invalid key: %w", err)
	}

	parts := strings.Split(string(rawKey), "|")
	if len(parts) != 2 {
		return coordinates.LatLng{}, coordinates.LatLng{}, fmt.Errorf("malformed key: %w", err)
	}

	home, err := coordinates.ParseLatLng(parts[0])
	if err != nil {
		return coordinates.LatLng{}, coordinates.LatLng{}, fmt.Errorf("parse home: %w", err)
	}

	work, err := coordinates.ParseLatLng(parts[1])
	if err != nil {
		return coordinates.LatLng{}, coordinates.LatLng{}, fmt.Errorf("parse work: %w", err)
	}

	return home, work, nil
}
