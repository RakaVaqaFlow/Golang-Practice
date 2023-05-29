package main

import (
	"client/internal/app"
	handler "client/internal/pkg/command_handler"
	"client/internal/pkg/tracer"
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
)

const defaultPort = "8080"

func main() {
	tp, err := tracer.NewTracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	ManagementSystemClient, err := app.NewClient(ctx, ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Start command handler.
	handler.CommandHandler(ctx, ManagementSystemClient)

}
