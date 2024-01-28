package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/strava/pkg/coordinates"
	"github.com/ViBiOh/strava/pkg/nominatim"
	"github.com/ViBiOh/strava/pkg/strava"
)

func main() {
	fs := flag.NewFlagSet("strava", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	stravaConfig := strava.Flags(fs, "")

	home := flags.New("Home", "Home LatLng").Prefix("strava").String(fs, "", nil)
	work := flags.New("Work", "Work LatLng").Prefix("strava").String(fs, "", nil)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(fmt.Errorf("config: %w", err))
	}

	logger.Init(&logger.Config{
		Level: "ERROR",
	})

	ctx := context.Background()

	homeLatLng, err := nominatim.GetLatLng(ctx, *home)
	FatalIfError(ctx, "get home lat lng", err)

	workLatLng, err := nominatim.GetLatLng(ctx, *work)
	FatalIfError(ctx, "get work lat lng", err)

	stravaApp, err := strava.New(ctx, stravaConfig)
	FatalIfError(ctx, "create strava", err)

	displayCommute(ctx, stravaApp, homeLatLng, workLatLng)
}

func displayCommute(ctx context.Context, stravaApp strava.App, homeLatLng, workLatLng coordinates.LatLng) {
	activities, err := stravaApp.GetActivities(ctx)
	FatalIfError(ctx, "get activities", err)

	roundTrips := make(map[string]uint8)

	for _, activity := range activities {
		if weekday := activity.StartDate.Weekday(); activity.Type != "Ride" || weekday < 0 || weekday > 5 {
			continue
		}

		day := activity.StartDate.Format(time.DateOnly)

		var found bool

		startLatLng, err := coordinates.NewLatLng(activity.StartLatlng)
		FatalIfError(ctx, "converting start", err)

		endLatLng, err := coordinates.NewLatLng(activity.EndLatlng)
		FatalIfError(ctx, "converting start", err)

		if startLatLng.IsWithin(homeLatLng, .5) {
			roundTrips[day] |= 1 << 3
			found = true
		}

		if endLatLng.IsWithin(workLatLng, .5) {
			roundTrips[day] |= 1 << 2
			found = true
		}

		if startLatLng.IsWithin(workLatLng, .5) {
			roundTrips[day] |= 1 << 1
			found = true
		}

		if endLatLng.IsWithin(homeLatLng, .5) {
			roundTrips[day] |= 1 << 0
			found = true
		}

		if !found {
			fmt.Println(activity.StartDate, activity.Name, "from", activity.StartLatlng, "to", activity.EndLatlng)
		}
	}

	output := make([]string, 0, len(roundTrips))

	for date, status := range roundTrips {
		item := fmt.Sprintf("%s | %04b", date, status)

		index := sort.Search(len(output), func(i int) bool {
			return output[i] > item
		})

		output = append(output, item)
		copy(output[index+1:], output[index:])
		output[index] = item
	}

	fmt.Printf("%s\n", strings.Join(output, "\n"))
}

func FatalIfError(ctx context.Context, label string, err error) {
	if err == nil {
		return
	}

	slog.LogAttrs(ctx, slog.LevelError, label, slog.Any("error", err))
	os.Exit(1)
}
