package main

import (
	"context"

	"github.com/ViBiOh/httputils/v4/pkg/logger"
)

func main() {
	config := newConfig()

	ctx := context.Background()

	newClients(ctx, config)

	services := newServices(config)

	err := services.commute.Start(ctx)
	logger.FatalfOnErr(ctx, err, "commute")
}
