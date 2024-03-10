package wahoo

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ViBiOh/commute/pkg/model"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

const (
	authURL   = "https://api.wahooligan.com/oauth/authorize"
	authToken = "https://api.wahooligan.com/oauth/token"
	apiURL    = "https://api.wahooligan.com/v1"
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

func (s Service) ID() string {
	return "wahoo"
}

func (s Service) LoginURL() string {
	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("response_type", "code")
	values.Add("scope", "workouts_read,plans_read")
	values.Add("redirect_uri", fmt.Sprintf("%s/token/%s", s.uri, s.ID()))

	return fmt.Sprintf("%s?%s", authURL, values.Encode())
}

func (s Service) ExchangeToken(ctx context.Context, r *http.Request) (string, error) {
	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("client_secret", s.clientSecret)
	values.Add("code", r.URL.Query().Get("code"))
	values.Add("grant_type", "authorization_code")
	values.Add("redirect_uri", fmt.Sprintf("%s/token/%s", s.uri, s.ID()))

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

func (s Service) Get(ctx context.Context, token string, before, after time.Time) (model.Rides, error) {
	requester := request.Get(apiURL).Header("Authorization", fmt.Sprintf("Bearer %s", token))

	workouts, err := s.getWorkouts(ctx, requester, before, after)
	if err != nil {
		return nil, fmt.Errorf("get workouts: %w", err)
	}

	rides, err := toRides(workouts)
	if err != nil {
		return nil, fmt.Errorf("get rides: %w", err)
	}

	return rides, nil
}

func (s Service) getWorkouts(ctx context.Context, requester request.Request, _, _ time.Time) ([]Workout, error) {
	params := url.Values{}
	params.Add("per_page", "100")

	resp, err := requester.Path("/workouts?%s", params.Encode()).Send(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	var workouts WorkoutsResponse
	if err = httpjson.Read(resp, &workouts); err != nil {
		return nil, err
	}

	return workouts.Workouts, nil
}

func toRides(_ []Workout) (model.Rides, error) {
	return nil, nil
}
