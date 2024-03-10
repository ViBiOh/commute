package commute

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ViBiOh/commute/pkg/coordinates"
	"github.com/ViBiOh/commute/pkg/mapbox"
	"github.com/ViBiOh/commute/pkg/model"
	"github.com/ViBiOh/commute/pkg/nominatim"
	"github.com/ViBiOh/commute/pkg/templ"
	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

const defaultDistance = 1.0

type Provider interface {
	ID() string
	ExchangeToken(ctx context.Context, r *http.Request) (string, error)
	Get(ctx context.Context, token string, before, after time.Time) (model.Rides, error)
}

type Service struct {
	providers map[string]Provider
	mapbox    mapbox.Service
	uri       string
}

func New(uri string, mapboxService mapbox.Service, providers ...Provider) Service {
	providersIndex := make(map[string]Provider)
	for _, provider := range providers {
		providersIndex[provider.ID()] = provider
	}

	return Service{
		providers: providersIndex,
		mapbox:    mapboxService,
	}
}

func (s Service) HandleToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		telemetry.SetRouteTag(ctx, "token")

		providerName := r.PathValue("provider")

		provider, ok := s.providers[providerName]
		if !ok {
			httperror.BadRequest(ctx, w, fmt.Errorf("unknown provider `%s`", providerName))
			return
		}

		token, err := provider.ExchangeToken(ctx, r)
		if err != nil {
			httperror.InternalServerError(ctx, w, err)
			return
		}

		rides, err := provider.Get(ctx, token, time.Time{}, time.Time{})
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

		templ.DisplayForm(ctx, w, s.uri, providerName, token, s.mapbox.StaticImage(clusters...), places)
	})
}

func (s Service) HandleCompute() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		telemetry.SetRouteTag(ctx, "compute")

		token := r.FormValue("token")
		providerName := r.FormValue("provider")

		provider, ok := s.providers[providerName]
		if !ok {
			httperror.BadRequest(ctx, w, fmt.Errorf("unknown provider `%s`", providerName))
			return
		}

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

		rides, err := provider.Get(ctx, token, before, after)
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
	})
}

func lastDayOfTheMonth(year, month int) time.Time {
	if month > 12 {
		month = 1
		year++
	}

	return time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
}
