package commute

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os/exec"
	"time"

	"github.com/ViBiOh/commute/pkg/model"
	"github.com/ViBiOh/commute/pkg/nominatim"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

const defaultDistance = 1.0

type Provider interface {
	LoginURL() string
	ExchangeToken(r *http.Request) (string, error)
	Get(ctx context.Context, token string, before, after time.Time) (model.Rides, error)
}

type Service struct {
	provider Provider
	server   *server.Server
	home     string
	work     string
}

type Config struct {
	Home string
	Work string
}

func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) *Config {
	var config Config

	flags.New("Home", "Home address").DocPrefix("commute").StringVar(fs, &config.Home, "", nil)
	flags.New("Work", "Work address").DocPrefix("commute").StringVar(fs, &config.Work, "", nil)

	return &config
}

func New(config *Config, server *server.Server, provider Provider) Service {
	return Service{
		provider: provider,
		server:   server,
		home:     config.Home,
		work:     config.Work,
	}
}

func (s Service) Start(ctx context.Context) error {
	homeCoord, _, err := nominatim.Geocode(ctx, s.home)
	if err != nil {
		return fmt.Errorf("get home: %w", err)
	}

	workCoord, _, err := nominatim.Geocode(ctx, s.work)
	if err != nil {
		return fmt.Errorf("get work: %w", err)
	}

	cmd := exec.Command("open", s.provider.LoginURL())
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("unable to open provider's url: %w", err)
	}

	var token string
	done := make(chan struct{})

	go s.server.Start(ctx, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/token" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		defer close(done)

		var err error

		token, err = s.provider.ExchangeToken(r)
		if err != nil {
			slog.Error(fmt.Sprintf("unable to exchange token: %s", err))
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Token exchanged with success, you can close this tab."))
	}))

	<-done
	s.server.Stop(ctx)

	rides, err := s.provider.Get(ctx, token, time.Time{}, time.Time{})
	if err != nil {
		return fmt.Errorf("get rides: %w", err)
	}

	commutes, err := rides.ToCommutes(homeCoord, workCoord, defaultDistance)
	if err != nil {
		return fmt.Errorf("to commutes: %w", err)
	}

	fmt.Println(commutes)

	return nil
}
