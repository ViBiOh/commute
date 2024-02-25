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

func (s Service) StravaLoginURL() string {
	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("response_type", "code")
	values.Add("scope", "read")
	values.Add("scope", "activity:read")
	values.Add("redirect_uri", fmt.Sprintf("%s/api/exchange_token", s.uri))

	return fmt.Sprintf("%s?%s", authURL, values.Encode())
}

func (s Service) handleStravaCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := s.exchangeToken(ctx, r)
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	rides, err := s.fetchRides(ctx, token)
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	getClusters(rides)

	templ.DisplayForm(ctx, w, s.uri, token, nil)
}

func (s Service) exchangeToken(ctx context.Context, r *http.Request) (string, error) {
	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("client_secret", s.clientSecret)
	values.Add("code", r.URL.Query().Get("code"))
	values.Add("grant+type", "authorization_code")

	resp, err := request.Post(authToken).Form(ctx, values)
	if err != nil {
		return "", fmt.Errorf("exchange token: %w", err)
	}

	var tokenResponse TokenResponse

	if err = httpjson.Read(resp, &tokenResponse); err != nil {
		return "", fmt.Errorf("read token: %w", err)
	}

	return tokenResponse.AccessToken, nil
}

func (s Service) handleCompute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token := r.FormValue("token")
	home, work, fields := s.geocodeAddresses(ctx, r)

	if fields.HasError() {
		templ.DisplayForm(ctx, w, s.uri, token, fields)
		return
	}

	rides, err := s.fetchRides(ctx, token)
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

func (s Service) fetchRides(ctx context.Context, token string) ([]Ride, error) {
	requester := request.Get(apiURL).Header("Authorization", fmt.Sprintf("Bearer %s", token))

	activities, err := s.getActivities(ctx, requester)
	if err != nil {
		return nil, fmt.Errorf("get activities: %w", err)
	}

	rides, err := getRides(activities)
	if err != nil {
		return nil, fmt.Errorf("get rides: %w", err)
	}

	return rides, nil
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
