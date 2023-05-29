package main

import (
	"context"
	"homework/internal/app/management_system/service"
	"homework/internal/pb"
	"homework/internal/pkg/db"
	dbname "homework/internal/pkg/repository/postgres"
	"homework/internal/tracer"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

const (
	serverPort = ":50051"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp, err := tracer.NewTracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	// database connection
	database, err := db.NewDB(ctx)
	if err != nil {
		return
	}
	defer database.GetPool(ctx).Close()
	tasks := dbname.NewTasks(database)
	users := dbname.NewUsers(database)

	// HTTP exporter для prometheus
	go http.ListenAndServe(":9091", promhttp.Handler())

	// start server
	server := grpc.NewServer()
	pb.RegisterManagementSystemSeviceServer(server, service.NewImplementation(users, tasks))

	lsn, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
