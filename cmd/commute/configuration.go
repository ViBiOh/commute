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
	"github.com/ViBiOh/httputils/v4/pkg/pprof"
	"github.com/ViBiOh/httputils/v4/pkg/server"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

type configuration struct {
	logger    *logger.Config
	alcotest  *alcotest.Config
	telemetry *telemetry.Config
	pprof     *pprof.Config
	health    *health.Config

	server *server.Config
	cors   *cors.Config
	owasp  *owasp.Config

	strava    *strava.Config
	wahoo     *wahoo.Config
	mapbox    *mapbox.Config
	publicURL *string
}

func newConfig() configuration {
	fs := flag.NewFlagSet("commute", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	config := configuration{
		logger:    logger.Flags(fs, "logger"),
		alcotest:  alcotest.Flags(fs, ""),
		telemetry: telemetry.Flags(fs, "telemetry"),
		pprof:     pprof.Flags(fs, "pprof"),
		health:    health.Flags(fs, ""),

		server: server.Flags(fs, ""),
		cors:   cors.Flags(fs, "cors"),
		owasp:  owasp.Flags(fs, "", flags.NewOverride("Csp", "default-src 'self'; base-uri 'self'; script-src 'self'; style-src 'self' 'httputils-nonce'; img-src 'self' api.mapbox.com/styles/v1/mapbox/dark-v11/")),

		strava:    strava.Flags(fs, "strava"),
		wahoo:     wahoo.Flags(fs, "wahoo"),
		mapbox:    mapbox.Flags(fs, "mapbox"),
		publicURL: flags.New("PublicURL", "Public URL for redirection").String(fs, "http://localhost:1080", nil),
	}

	_ = fs.Parse(os.Args[1:])

	return config
}
