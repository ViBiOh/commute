package strava

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/strava/pkg/coordinates"
	"github.com/ViBiOh/strava/pkg/nominatim"
	"github.com/ViBiOh/strava/pkg/templ"
)

func (s Service) Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		action := r.PathValue("action")
		telemetry.SetRouteTag(ctx, action)

		switch action {
		case "compute":
			s.handleCompute(w, r)

		case "exchange_token":
			s.handleStravaCallback(w, r)

		default:
			httperror.NotFound(ctx, w)
		}
	})
}

func (s Service) handleCompute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	home, work, fields := s.geocodeAddresses(ctx, r)

	if fields.HasError() {
		templ.DisplayForm(ctx, w, s.uri, fields)
		return
	}

	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("response_type", "code")
	values.Add("scope", "read")
	values.Add("scope", "activity:read")
	values.Add("redirect_uri", fmt.Sprintf("%s/api/exchange_token?%s", s.uri, encodeKey(home, work)))

	http.Redirect(w, r, fmt.Sprintf("%s?%s", authURL, values.Encode()), http.StatusFound)
}

func (s Service) geocodeAddresses(ctx context.Context, r *http.Request) (coordinates.LatLng, coordinates.LatLng, templ.Fields) {
	fields := templ.Fields{}

	var home, work coordinates.LatLng

	home, fields["home"] = s.geocodeAddress(ctx, r.FormValue("home"))
	work, fields["work"] = s.geocodeAddress(ctx, r.FormValue("work"))

	return home, work, fields
}

func (s Service) geocodeAddress(ctx context.Context, value string) (coordinates.LatLng, templ.Field) {
	field := templ.Field{
		Value: value,
	}

	var latLng coordinates.LatLng

	latLng, field.Value, field.Err = nominatim.GetLatLng(ctx, value)

	return latLng, field
}

func (s Service) handleStravaCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	home, work, err := parseKey(r)
	if err != nil {
		httperror.BadRequest(ctx, w, fmt.Errorf("parse key: %w", err))
		return
	}

	requester, err := s.exchangeToken(ctx, r)
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	activities, err := s.getActivities(ctx, requester)
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	rides, err := getRides(activities)
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	commutes, err := getCommutes(rides, home, work)
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	templ.DisplayResult(ctx, w, s.uri, s.mapboxToken, home, work, commutes)
}

func (s Service) exchangeToken(ctx context.Context, r *http.Request) (request.Request, error) {
	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("client_secret", s.clientSecret)
	values.Add("code", r.URL.Query().Get("code"))
	values.Add("grant+type", "authorization_code")

	resp, err := request.Post(authToken).Form(ctx, values)
	if err != nil {
		return request.Request{}, fmt.Errorf("exchange token: %w", err)
	}

	var tokenResponse TokenResponse

	if err = httpjson.Read(resp, &tokenResponse); err != nil {
		return request.Request{}, fmt.Errorf("read token: %w", err)
	}

	return request.Get(apiURL).Header("Authorization", fmt.Sprintf("Bearer %s", tokenResponse.AccessToken)), nil
}
