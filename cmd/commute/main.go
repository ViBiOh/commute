package main

import (
	"context"

	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/recoverer"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

func main() {
	configs := newConfig()
	alcotest.DoAndExit(configs.alcotest)

	ctx := context.Background()

	clients, err := newClient(ctx, configs)
	logger.FatalfOnErr(ctx, err, "client")

	defer clients.Close(ctx)
	go clients.Start()

	service := newService(configs)

	httpServer := server.New(configs.http)

	go httpServer.Start(clients.health.EndCtx(), httputils.Handler(
		newPort(configs, service),
		clients.health,
		recoverer.Middleware,
		clients.telemetry.Middleware("http"),
		owasp.New(configs.owasp).Middleware,
		cors.New(configs.cors).Middleware,
	))

	clients.health.WaitForTermination(httpServer.Done())

	server.GracefulWait(httpServer.Done())
}
