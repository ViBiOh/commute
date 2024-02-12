package main

import (
	"net/http"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/strava/pkg/templ"
)

const (
	title         = "Calcul Indemnites Kilométriques Vélo"
	exchangeToken = "/api/exchange_token"
)

func newPort(config configuration, service service) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(exchangeToken, http.StripPrefix(exchangeToken, service.strava.Handle()))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonce := owasp.Nonce()
		owasp.WriteNonce(w, nonce)

		ctx := r.Context()

		component := templ.Hello(nonce, title, "world")
		if err := component.Render(ctx, w); err != nil {
			httperror.InternalServerError(ctx, w, err)
		}
	}))

	return mux
}
