package main

import (
	"context"

	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

func main() {
	configs := newConfig()
	alcotest.DoAndExit(configs.alcotest)

	ctx := context.Background()

	clients, err := newClients(ctx, configs)
	logger.FatalfOnErr(ctx, err, "client")

	defer clients.Close(ctx)
	go clients.Start()

	service := newServices(configs)

	httpServer := server.New(configs.server)

	go httpServer.Start(clients.health.EndCtx(), httputils.Handler(
		newPort(configs, service),
		clients.health,
		clients.telemetry.Middleware("http"),
		owasp.New(configs.owasp).Middleware,
		cors.New(configs.cors).Middleware,
	))

	clients.health.WaitForTermination(httpServer.Done())

	server.GracefulWait(httpServer.Done())
}
