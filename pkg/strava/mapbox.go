package strava

import (
	"fmt"
	"strings"

	"github.com/ViBiOh/strava/pkg/coordinates"
)

var colors = []string{
	"ff2600", // red
	"0433ff", // blue
	"ff9300", // orange
	"ff40ff", // purple
	"919191", // grey
	"aa7942", // brown
}

func (s Service) getMapboxStaticImage(coords ...coordinates.LatLng) string {
	var builder strings.Builder
	builder.WriteString("https://api.mapbox.com/styles/v1/mapbox/dark-v11/static/")

	for index, coord := range coords {
		if index != 0 {
			builder.WriteString(",")
		}

		builder.WriteString(fmt.Sprintf("pin-l-%d+%s(%s)", index+1, colors[index%len(colors)], coord.LngLat()))
	}

	builder.WriteString("/auto/400x300@2x?access_token=")
	builder.WriteString(s.mapboxToken)

	return builder.String()
}
