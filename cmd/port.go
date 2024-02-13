package main

import (
	"net/http"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/strava/pkg/templ"
)

const (
	title   = "Compute Indemnites Kilométriques Vélo"
	apiPath = "/api"
)

func newPort(config configuration, service service) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(apiPath+"/compute", http.StripPrefix(apiPath, service.strava.Handle()))
	mux.Handle(apiPath+"/exchange_token", http.StripPrefix(apiPath, service.strava.Handle()))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonce := owasp.Nonce()
		owasp.WriteNonce(w, nonce)

		ctx := r.Context()

		component := templ.Form(nonce, title)
		if err := component.Render(ctx, w); err != nil {
			httperror.InternalServerError(ctx, w, err)
		}
	}))

	return mux
}
