package templ

import (
	"context"
	"net/http"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/strava/pkg/coordinates"
	"github.com/ViBiOh/strava/pkg/model"
)

const title = "Compute Indemnites Kilométriques Vélo"

func Handler(uri string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DisplayForm(r.Context(), w, uri, nil)
	})
}

func DisplayForm(ctx context.Context, w http.ResponseWriter, uri string, fields Fields) {
	nonce := owasp.Nonce()
	owasp.WriteNonce(w, nonce)

	telemetry.SetRouteTag(ctx, "index")

	component := Form(uri, nonce, title, fields)
	if err := component.Render(ctx, w); err != nil {
		httperror.InternalServerError(ctx, w, err)
	}
}

func DisplayResult(ctx context.Context, w http.ResponseWriter, uri string, home, work coordinates.LatLng, commutes model.Commutes) {
	nonce := owasp.Nonce()
	owasp.WriteNonce(w, nonce)

	telemetry.SetRouteTag(ctx, "result")

	component := Result(uri, nonce, title, home, work, commutes)
	if err := component.Render(ctx, w); err != nil {
		httperror.InternalServerError(ctx, w, err)
	}
}
