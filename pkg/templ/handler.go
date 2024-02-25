package templ

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/strava/pkg/coordinates"
	"github.com/ViBiOh/strava/pkg/model"
)

const title = "Compute Indemnites Kilométriques Vélo"

func Handler(uri, loginURL string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderLogin(r.Context(), w, uri, loginURL)
	})
}

func RenderLogin(ctx context.Context, w http.ResponseWriter, uri, loginURL string) {
	nonce := owasp.Nonce()
	owasp.WriteNonce(w, nonce)

	telemetry.SetRouteTag(ctx, "login")

	component := Login(uri, nonce, title, loginURL)
	if err := component.Render(ctx, w); err != nil {
		httperror.InternalServerError(ctx, w, err)
	}
}

func DisplayForm(ctx context.Context, w http.ResponseWriter, uri, token string, fields Fields) {
	nonce := owasp.Nonce()
	owasp.WriteNonce(w, nonce)

	telemetry.SetRouteTag(ctx, "form")

	component := Form(uri, nonce, title, token, fields)
	if err := component.Render(ctx, w); err != nil {
		httperror.InternalServerError(ctx, w, err)
	}
}

func DisplayResult(ctx context.Context, w http.ResponseWriter, uri string, mapboxToken string, home, work coordinates.LatLng, commutes model.Commutes) {
	nonce := owasp.Nonce()
	owasp.WriteNonce(w, nonce)

	telemetry.SetRouteTag(ctx, "result")

	staticValues := url.Values{}
	staticValues.Add("access_token", mapboxToken)

	staticImageURL := fmt.Sprintf("https://api.mapbox.com/styles/v1/mapbox/dark-v11/static/pin-s-home+669c35(%s),pin-s+0061ff(%s)/auto/400x300@2x?%s", home.LngLat(), work.LngLat(), staticValues.Encode())

	component := Result(uri, nonce, title, staticImageURL, commutes)
	if err := component.Render(ctx, w); err != nil {
		httperror.InternalServerError(ctx, w, err)
	}
}
