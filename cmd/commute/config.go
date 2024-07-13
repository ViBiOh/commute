package main

import (
	"flag"
	"os"

	"github.com/ViBiOh/commute/pkg/commute"
	"github.com/ViBiOh/commute/pkg/strava"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

type configuration struct {
	logger *logger.Config
	health *health.Config

	server *server.Config

	commute *commute.Config
	strava  *strava.Config
}

func newConfig() configuration {
	fs := flag.NewFlagSet("commute", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	config := configuration{
		logger: logger.Flags(fs, "logger"),
		health: health.Flags(fs, ""),

		server: server.Flags(fs, ""),

		commute: commute.Flags(fs, "commute"),
		strava:  strava.Flags(fs, "strava"),
	}

	_ = fs.Parse(os.Args[1:])

	return config
}
