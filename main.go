package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/strava/pkg/strava"
)

func main() {
	var homeLatLng []float64
	var workLatLng []float64

	fs := flag.NewFlagSet("strava", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	stravaConfig := strava.Flags(fs, "")

	flags.New("Home", "Home LatLng").Prefix("strava").Float64SliceVar(fs, &homeLatLng, nil, nil)
	flags.New("Work", "Work LatLng").Prefix("strava").Float64SliceVar(fs, &workLatLng, nil, nil)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(fmt.Errorf("config: %w", err))
	}

	logger.Init(&logger.Config{
		Level: "ERROR",
	})

	ctx := context.Background()

	stravaApp, err := strava.New(ctx, stravaConfig)
	if err != nil {
		slog.ErrorContext(ctx, "create strava", "error", err)
		return
	}

	displayCommute(ctx, stravaApp, homeLatLng, workLatLng)
}

func displayCommute(ctx context.Context, stravaApp strava.App, homeLatLng, workLatLng []float64) {
	activities, err := stravaApp.GetActivities(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "get activities", "error", err)
		return
	}

	roundTrip := make(map[string]uint8)

	for _, activity := range activities {
		if weekday := activity.StartDate.Weekday(); activity.Type != "Ride" || weekday < 0 || weekday > 5 {
			continue
		}

		day := activity.StartDate.Format(time.DateOnly)
		var found bool

		if slices.Equal(activity.StartLatlng, homeLatLng) {
			roundTrip[day] |= 1 << 3
			found = true
		}

		if slices.Equal(activity.EndLatlng, workLatLng) {
			roundTrip[day] |= 1 << 2
			found = true
		}

		if slices.Equal(activity.StartLatlng, workLatLng) {
			roundTrip[day] |= 1 << 1
			found = true
		}

		if slices.Equal(activity.EndLatlng, homeLatLng) {
			roundTrip[day] |= 1 << 0
			found = true
		}

		if !found {
			fmt.Println(activity.StartDate, activity.Name, "from", activity.StartLatlng, "to", activity.EndLatlng)
		}
	}

	output := make([]string, 0, len(roundTrip))

	for date, status := range roundTrip {
		output = append(output, fmt.Sprintf("%s | %04b", date, status))
	}

	sort.Strings(output)

	fmt.Printf("%s\n", strings.Join(output, "\n"))
}
