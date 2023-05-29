package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
	"workshop-8-3/api/internal"
	"workshop-8-3/api/internal/app/todo/pb"
	"workshop-8-3/api/internal/config"
	"workshop-8-3/api/internal/database"
	"workshop-8-3/api/internal/database/todo"
	"workshop-8-3/api/internal/model"
)

type Implementation struct {
	pb.UnimplementedTodoServiceServer

	db database.PGX
}

func NewImplementation(db database.PGX) *Implementation {
	return &Implementation{
		db: db,
	}
}

func (s *Implementation) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	tr := otel.Tracer("DeleteTodo")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(in.String()))
	defer span.End()

	ok, err := todo.NewDeleter().Delete(ctx, s.db, int(in.Id))
	if err != nil {
		return nil, err
	}
	internal.DeletedCounter.Add(1)

	return &pb.DeleteTodoResponse{Ok: ok}, nil
}

func (s *Implementation) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	tr := otel.Tracer("CreateTodo")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(in.String()))
	defer span.End()

	id, err := todo.NewCreator().Create(ctx, s.db, model.CreateInput{
		UserID: in.UserId,
		Text:   in.Text,
	})
	if err != nil {
		return nil, err
	}

	internal.RegCounter.Add(1)

	return &pb.CreateTodoResponse{
		Id: uint32(id),
	}, nil
}

func (s *Implementation) ListTodo(ctx context.Context, in *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	todos, err := todo.NewSearcher().Search(ctx, s.db, model.Pagination{
		Page:  int(in.Pagination.Page),
		Limit: int(in.Pagination.Limit),
	})
	if err != nil {
		return nil, err
	}
	result := make([]*pb.Todo, 0, len(todos))
	for _, m := range todos {
		result = append(result, &pb.Todo{
			Id:   uint32(m.ID),
			Text: m.Text,
		})
	}
	return &pb.ListTodoResponse{
		Todos: result,
	}, nil
}

const (
	service     = "api"
	environment = "development"
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
		)),
	)
	return tp, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ошибка при чтении env файла: %v", err)
	}
	cfg := config.Config{}
	if err = env.Parse(&cfg); err != nil {
		log.Fatalf("не могу распарсить конфиг: %v", err)
	}

	tp, err := tracerProvider("http://localhost:14268/api/traces")
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

	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDbHost, cfg.PostgresDb),
	)
	if err != nil {
		log.Fatalf("невозможно подключиться к базе: %v", err)
	}

	err = goose.Up(db.DB, "./migrations")
	if err != nil {
		log.Fatalf("невозможно накатить миграции: %v", err)
	}

	// HTTP exporter для prometheus
	go http.ListenAndServe(":9091", promhttp.Handler())

	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(server, NewImplementation(database.NewPGX(db)))

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
