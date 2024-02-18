package strava

import (
	"flag"

	"github.com/ViBiOh/flags"
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

	flags.New("ClientID", "App Client ID").DocPrefix("strava").StringVar(fs, &config.ClientID, "", nil)
	flags.New("ClientSecret", "App Client Secret").DocPrefix("strava").StringVar(fs, &config.ClientSecret, "", nil)

	return &config
}

func New(config *Config, uri string) Service {
	return Service{
		uri:          uri,
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
	}
}
