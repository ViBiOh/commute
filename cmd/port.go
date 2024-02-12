package main

import (
	"net/http"
)

const exchangeToken = "/api/exchange_token"

func newPort(config configuration, service service) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(exchangeToken, http.StripPrefix(exchangeToken, service.strava.Handle()))

	return mux
}
