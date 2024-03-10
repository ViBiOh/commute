package main

import (
	"flag"
	"os"

	"github.com/ViBiOh/commute/pkg/mapbox"
	"github.com/ViBiOh/commute/pkg/strava"
	"github.com/ViBiOh/commute/pkg/wahoo"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/server"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

type configuration struct {
	alcotest  *alcotest.Config
	telemetry *telemetry.Config
	logger    *logger.Config
	cors      *cors.Config
	owasp     *owasp.Config
	http      *server.Config
	health    *health.Config

	strava    *strava.Config
	wahoo     *wahoo.Config
	mapbox    *mapbox.Config
	publicURL *string
}

func newConfig() (configuration, error) {
	fs := flag.NewFlagSet("commute", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	return configuration{
		http:      server.Flags(fs, ""),
		health:    health.Flags(fs, ""),
		alcotest:  alcotest.Flags(fs, ""),
		logger:    logger.Flags(fs, "logger"),
		telemetry: telemetry.Flags(fs, "telemetry"),
		owasp:     owasp.Flags(fs, "", flags.NewOverride("Csp", "default-src 'self'; base-uri 'self'; script-src 'self'; style-src 'self' 'httputils-nonce'; img-src 'self' api.mapbox.com/styles/v1/mapbox/dark-v11/")),
		cors:      cors.Flags(fs, "cors"),

		strava:    strava.Flags(fs, "strava"),
		wahoo:     wahoo.Flags(fs, "wahoo"),
		mapbox:    mapbox.Flags(fs, "mapbox"),
		publicURL: flags.New("PublicURL", "Public URL for redirection").String(fs, "http://localhost:1080", nil),
	}, fs.Parse(os.Args[1:])
}
