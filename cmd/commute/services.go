package main

import (
	"strconv"

	"github.com/ViBiOh/commute/pkg/commute"
	"github.com/ViBiOh/commute/pkg/strava"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

type services struct {
	server  *server.Server
	strava  strava.Service
	commute commute.Service
}

func newServices(config configuration) services {
	var output services

	output.server = server.New(config.server)

	output.strava = strava.New(config.strava, getListenAddr(config.server))
	output.commute = commute.New(config.commute, output.server, output.strava)

	return output
}

func getListenAddr(config *server.Config) string {
	port := config.Port
	if port == 0 {
		return ""
	}

	address := config.Address
	if len(address) == 0 {
		address = "127.0.0.1"
	}

	protocol := "http://"
	if len(config.Cert) != 0 && len(config.Key) != 0 {
		protocol += "s"
	}

	return protocol + address + ":" + strconv.FormatUint(uint64(port), 10)
}
