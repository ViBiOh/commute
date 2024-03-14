package mapbox

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ViBiOh/commute/pkg/coordinates"
	"github.com/ViBiOh/flags"
)

var colors = []string{
	"ff2600", // red
	"0433ff", // blue
	"ff9300", // orange
	"ff40ff", // purple
	"919191", // grey
	"aa7942", // brown
}

type Service struct {
	token string
}

type Config struct {
	Token string
}

func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) *Config {
	var config Config

	flags.New("AccessToken", "Mapbox Access Token").Prefix(prefix).DocPrefix("mapbox").StringVar(fs, &config.Token, "", nil)

	return &config
}

func New(config *Config) Service {
	return Service{
		token: config.Token,
	}
}

func (s Service) StaticImage(coords ...coordinates.LatLng) []string {
	var builder strings.Builder
	builder.WriteString("https://api.mapbox.com/styles/v1/mapbox/dark-v11/static/")

	for index, coord := range coords {
		if index != 0 {
			builder.WriteString(",")
		}

		builder.WriteString(fmt.Sprintf("pin-l-%d+%s(%s)", index+1, colors[index%len(colors)], coord.LngLat()))
	}

	builder.WriteString("/auto/800x600")

	url := builder.String()

	return []string{
		url + "?access_token=" + s.token,
		url + "@2x?access_token=" + s.token + " 2x",
	}
}
