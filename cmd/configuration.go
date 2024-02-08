package main

import (
	"flag"
	"os"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/server"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/strava/pkg/strava"
)

type configuration struct {
	alcotest  *alcotest.Config
	telemetry *telemetry.Config
	logger    *logger.Config
	cors      *cors.Config
	owasp     *owasp.Config
	http      *server.Config
	health    *health.Config

	strava *strava.Config
	home   *string
	work   *string
}

func newConfig() (configuration, error) {
	fs := flag.NewFlagSet("strava", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	return configuration{
		http:      server.Flags(fs, ""),
		health:    health.Flags(fs, ""),
		alcotest:  alcotest.Flags(fs, ""),
		logger:    logger.Flags(fs, "logger"),
		telemetry: telemetry.Flags(fs, "telemetry"),
		owasp:     owasp.Flags(fs, ""),
		cors:      cors.Flags(fs, "cors"),

		strava: strava.Flags(fs, ""),

		home: flags.New("Home", "Home LatLng").DocPrefix("strava").String(fs, "Address of your Home", nil),
		work: flags.New("Work", "Work LatLng").DocPrefix("strava").String(fs, "Address of your Work", nil),
	}, fs.Parse(os.Args[1:])
}
