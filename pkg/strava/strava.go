package strava

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ViBiOh/commute/pkg/coordinates"
	"github.com/ViBiOh/commute/pkg/model"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

const (
	authURL   = "https://www.strava.com/oauth/authorize"
	authToken = "https://www.strava.com/oauth/token"
	apiURL    = "https://www.strava.com/api/v3"
)

type Service struct {
	uri          string
	clientID     string
	clientSecret string
}

type Config struct {
	ClientID     string
	ClientSecret string
}

func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) *Config {
	var config Config

	flags.New("ClientID", "App Client ID").Prefix(prefix).DocPrefix("strava").StringVar(fs, &config.ClientID, "", nil)
	flags.New("ClientSecret", "App Client Secret").Prefix(prefix).DocPrefix("strava").StringVar(fs, &config.ClientSecret, "", nil)

	return &config
}

func New(config *Config, uri string) Service {
	return Service{
		uri:          uri,
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
	}
}

func (s Service) LoginURL() string {
	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("response_type", "code")
	values.Add("scope", "read")
	values.Add("scope", "activity:read")
	values.Add("redirect_uri", fmt.Sprintf("%s/token", s.uri))

	return fmt.Sprintf("%s?%s", authURL, values.Encode())
}

func (s Service) ExchangeToken(r *http.Request) (string, error) {
	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("client_secret", s.clientSecret)
	values.Add("code", r.URL.Query().Get("code"))
	values.Add("grant+type", "authorization_code")

	resp, err := request.Post(authToken).Form(r.Context(), values)
	if err != nil {
		return "", fmt.Errorf("exchange token: %w", err)
	}

	tokenResponse, err := httpjson.Read[TokenResponse](resp)
	if err != nil {
		return "", fmt.Errorf("read token: %w", err)
	}

	return tokenResponse.AccessToken, nil
}

func (s Service) Get(ctx context.Context, token string, before, after time.Time) (model.Rides, error) {
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

func (s Service) getActivities(ctx context.Context, requester request.Request, before, after time.Time) ([]Activity, error) {
	params := url.Values{}
	params.Add("per_page", "100")

	if !before.IsZero() {
		params.Add("before", strconv.FormatInt(before.Unix(), 10))
	}

	if !after.IsZero() {
		params.Add("after", strconv.FormatInt(after.Unix(), 10))
	}

	resp, err := requester.Path("/athlete/activities?%s", params.Encode()).Send(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return httpjson.Read[[]Activity](resp)
}

func toRides(activities []Activity) (model.Rides, error) {
	var output []model.Ride

	for _, activity := range activities {
		if activity.Type != "Ride" {
			continue
		}

		start, err := coordinates.NewLatLng(activity.StartLatlng)
		if err != nil && !activity.Commute {
			return nil, fmt.Errorf("parse start: %w", err)
		}

		end, err := coordinates.NewLatLng(activity.EndLatlng)
		if err != nil && !activity.Commute {
			return nil, fmt.Errorf("parse end: %w", err)
		}

		output = append(output, model.Ride{
			ID:       strconv.Itoa(activity.ID),
			Date:     activity.StartDate,
			Duration: time.Duration(activity.ElapsedTime) * time.Second,
			Start:    start,
			End:      end,
			Commute:  activity.Commute,
		})
	}

	return output, nil
}
