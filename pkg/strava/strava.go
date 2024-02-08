package strava

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"sort"
	"strings"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/strava/pkg/coordinates"
)

const (
	authURL   = "https://www.strava.com/oauth/authorize"
	authToken = "https://www.strava.com/oauth/token"
	apiURL    = "https://www.strava.com/api/v3"
)

type Service struct {
	clientID     string
	clientSecret string
}

type Config struct {
	ClientID     string
	ClientSecret string
}

func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) *Config {
	var config Config

	flags.New("ClientID", "App Client ID").DocPrefix("strava").StringVar(fs, &config.ClientID, "", nil)
	flags.New("ClientSecret", "App Client Secret").DocPrefix("strava").StringVar(fs, &config.ClientSecret, "", nil)

	return &config
}

func New(config *Config) Service {
	return Service{
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
	}
}

func (s Service) Open(uri string, home, work coordinates.LatLng) error {
	redirectValues := url.Values{}
	redirectValues.Add("key", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s|%s", home.String(), work.String()))))

	values := url.Values{}
	values.Add("client_id", s.clientID)
	values.Add("response_type", "code")
	values.Add("approval_prompt", "force")
	values.Add("scope", "read")
	values.Add("scope", "activity:read")
	values.Add("redirect_uri", fmt.Sprintf("%s/api/exchange_token?%s", uri, redirectValues.Encode()))

	loginURL := fmt.Sprintf("%s?%s", authURL, values.Encode())

	return exec.Command("open", loginURL).Run()
}

func (s Service) Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		commutes, err := computeCommute(activities, home, work)
		if err != nil {
			httperror.InternalServerError(ctx, w, err)
			return
		}

		formatCommute(commutes, w)
	})
}

func formatCommute(commutes map[string]uint8, w io.Writer) {
	output := make([]string, 0, len(commutes))

	for date, status := range commutes {
		item := fmt.Sprintf("%s | %04b", date, status)

		index := sort.Search(len(output), func(i int) bool {
			return output[i] > item
		})

		output = append(output, item)
		copy(output[index+1:], output[index:])
		output[index] = item
	}

	fmt.Fprintf(w, "%s\n", strings.Join(output, "\n"))
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

func parseKey(r *http.Request) (coordinates.LatLng, coordinates.LatLng, error) {
	rawKey, err := base64.StdEncoding.DecodeString(r.URL.Query().Get("key"))
	if err != nil {
		return coordinates.LatLng{}, coordinates.LatLng{}, fmt.Errorf("invalid key: %w", err)
	}

	parts := strings.Split(string(rawKey), "|")
	if len(parts) != 2 {
		return coordinates.LatLng{}, coordinates.LatLng{}, fmt.Errorf("malformed key: %w", err)
	}

	home, err := coordinates.ParseLatLng(parts[0])
	if err != nil {
		return coordinates.LatLng{}, coordinates.LatLng{}, fmt.Errorf("parse home: %w", err)
	}

	work, err := coordinates.ParseLatLng(parts[1])
	if err != nil {
		return coordinates.LatLng{}, coordinates.LatLng{}, fmt.Errorf("parse work: %w", err)
	}

	return home, work, nil
}
