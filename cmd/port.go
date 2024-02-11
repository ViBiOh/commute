package main

import (
	"net/http"

	"github.com/ViBiOh/httputils/v4/pkg/renderer"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

const exchangeToken = "/api/exchange_token"

func newPort(config configuration, adapter adapter, service service) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(exchangeToken, http.StripPrefix(exchangeToken, service.strava.Handle()))

	mux.Handle("/", adapter.renderer.Handler(func(w http.ResponseWriter, r *http.Request) (renderer.Page, error) {
		telemetry.SetRouteTag(r.Context(), "index")

		return renderer.NewPage("public", http.StatusOK, nil), nil
	}))

	return mux
}
