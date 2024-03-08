package strava

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/strava/pkg/coordinates"
	"github.com/ViBiOh/strava/pkg/model"
	"github.com/ViBiOh/strava/pkg/nominatim"
	"github.com/ViBiOh/strava/pkg/templ"
)

const defaultDistance = 1.0

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

	rides, err := s.fetchRides(ctx, token, time.Time{}, time.Time{})
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	clusters := rides.Coordinates().Clusters(defaultDistance)

	places := make([]templ.Place, len(clusters))

	for index, cluster := range clusters {
		name, err := nominatim.Reverse(ctx, cluster)
		if err != nil {
			httperror.InternalServerError(ctx, w, fmt.Errorf("reverse geocode `%f`: %w", cluster, err))
			return
		}

		places[index] = templ.Place{
			Coordinates: cluster,
			Name:        name,
		}
	}

	templ.DisplayForm(ctx, w, s.uri, token, s.mapbox.StaticImage(clusters...), places)
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

	month, err := strconv.Atoi(r.FormValue("month"))
	if err != nil {
		httperror.BadRequest(ctx, w, fmt.Errorf("parse month: %w", err))
		return
	}

	home, err := coordinates.ParseLatLng(r.FormValue("home"))
	if err != nil {
		httperror.BadRequest(ctx, w, fmt.Errorf("parse home: %w", err))
		return
	}

	work, err := coordinates.ParseLatLng(r.FormValue("work"))
	if err != nil {
		httperror.BadRequest(ctx, w, fmt.Errorf("parse work: %w", err))
		return
	}

	now := time.Now()
	year := now.Year()
	if month > int(now.Month()) {
		year--
	}

	after := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	before := lastDayOfTheMonth(year, month+1)

	rides, err := s.fetchRides(ctx, token, before, after)
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	commutes, err := rides.ToCommutes(home, work, defaultDistance)
	if err != nil {
		httperror.InternalServerError(ctx, w, err)
		return
	}

	templ.DisplayResult(ctx, w, s.uri, s.mapbox.StaticImage(home, work), home, work, commutes)
}

func (s Service) fetchRides(ctx context.Context, token string, before, after time.Time) (model.Rides, error) {
	requester := request.Get(apiURL).Header("Authorization", fmt.Sprintf("Bearer %s", token))

	activities, err := s.getActivities(ctx, requester, before, after)
	if err != nil {
		return nil, fmt.Errorf("get activities: %w", err)
	}

	rides, err := toRides(activities)
	if err != nil {
		return nil, fmt.Errorf("get rides: %w", err)
	}

	return rides, nil
}

func lastDayOfTheMonth(year, month int) time.Time {
	if month > 12 {
		month = 1
		year++
	}

	return time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
}
