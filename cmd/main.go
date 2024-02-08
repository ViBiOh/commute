package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/recoverer"
	"github.com/ViBiOh/httputils/v4/pkg/server"
	"github.com/ViBiOh/strava/pkg/nominatim"
)

func main() {
	config, err := newConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("config: %s", err))
	}

	alcotest.DoAndExit(config.alcotest)

	go func() {
		fmt.Println(http.ListenAndServe("localhost:9999", http.DefaultServeMux))
	}()

	ctx := context.Background()

	client, err := newClient(ctx, config)
	logger.FatalfOnErr(ctx, err, "client")

	service := newService(config)

	defer client.Close(ctx)

	httpServer := server.New(config.http)

	go httpServer.Start(client.health.EndCtx(), httputils.Handler(newPort(config, service), client.health, recoverer.Middleware, client.telemetry.Middleware("http"), owasp.New(config.owasp).Middleware, cors.New(config.cors).Middleware))

	home, err := nominatim.GetLatLng(ctx, *config.home)
	logger.FatalfOnErr(ctx, err, "home")

	work, err := nominatim.GetLatLng(ctx, *config.work)
	logger.FatalfOnErr(ctx, err, "work")

	err = service.strava.Open(fmt.Sprintf("http://localhost:%d", config.http.Port), home, work)
	logger.FatalfOnErr(ctx, err, "open")

	client.health.WaitForTermination(httpServer.Done())

	server.GracefulWait(httpServer.Done())
}
