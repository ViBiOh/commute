package strava

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
	"github.com/ViBiOh/httputils/v4/pkg/recoverer"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

const (
	authURL   = "https://www.strava.com/oauth/authorize"
	authToken = "https://www.strava.com/oauth/token"
	apiURL    = "https://www.strava.com/api/v3"
)

type App struct {
	ClientID string
	request  request.Request
}

type StravaConfig struct {
	server       *server.Config
	ClientID     string
	ClientSecret string
}

func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) *StravaConfig {
	var config StravaConfig

	flags.New("ClientID", "App Client ID").DocPrefix("strava").StringVar(fs, &config.ClientID, "", nil)
	flags.New("ClientSecret", "App Client Secret").DocPrefix("strava").StringVar(fs, &config.ClientSecret, "", nil)

	config.server = server.Flags(fs, "")

	return &config
}

func New(ctx context.Context, config *StravaConfig) (App, error) {
	loginURL := fmt.Sprintf("%s?client_id=%s&response_type=code&redirect_uri=http://127.0.0.1:%d/exchange_token&approval_prompt=force&scope=read&scope=activity:read", authURL, config.ClientID, config.server.Port)
	if err := exec.Command("open", loginURL).Run(); err != nil {
		return App{}, fmt.Errorf("open login URL: %w", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var token string

	oauthServer := server.New(config.server)
	oauthServer.Start(ctx, httputils.Handler(oauthMux(config.ClientID, config.ClientSecret, cancel, &token), nil, recoverer.Middleware))

	return App{
		ClientID: config.ClientID,
		request:  request.Get(apiURL).Header("Authorization", fmt.Sprintf("Bearer %s", token)),
	}, nil
}

func oauthMux(clientID, clientSecret string, cancel func(), token *string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/exchange_token", func(w http.ResponseWriter, r *http.Request) {
		values := url.Values{}
		values.Add("client_id", clientID)
		values.Add("client_secret", clientSecret)
		values.Add("code", r.URL.Query().Get("code"))
		values.Add("grant+type", "authorization_code")

		ctx := r.Context()

		resp, err := request.Post(authToken).Form(ctx, values)
		if err != nil {
			httperror.InternalServerError(ctx, w, fmt.Errorf("exchange token: %w", err))
			return
		}

		var tokenResponse TokenResponse

		if err = httpjson.Read(resp, &tokenResponse); err != nil {
			httperror.InternalServerError(ctx, w, err)
			return
		}

		*token = tokenResponse.AccessToken

		_, _ = w.Write([]byte(`Authentication succeeded, you can close this tab.`))

		cancel()
	})

	return mux
}

func (a App) GetActivities(ctx context.Context) ([]Activity, error) {
	resp, err := a.request.Path("/athlete/activities?per_page=100").Send(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("get activities: %w", err)
	}

	var activities []Activity
	if err = httpjson.Read(resp, &activities); err != nil {
		return nil, err
	}

	return activities, nil
}
