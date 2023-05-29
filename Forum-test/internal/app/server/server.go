package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/storm5758/Forum-test/pkg/api"
	gw_api "github.com/storm5758/Forum-test/pkg/gw/api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	ServerAdressGRPC = ":6000"
	ServerAdressHTTP = ":5000"
	SwaggerDir       = "./swagger"

	GRPCTimeoutConnection = 5 * time.Second
)

type Services struct {
	User   api.UserServer
	Forum  api.ForumServer
	Thread api.ThreadServer
	Post   api.PostServer
	Admin  api.AdminServer
}

type closer func() error

type server struct {
	Services
	lis        net.Listener
	grpcServer *grpc.Server
	group      errgroup.Group
	closers    []closer
}

func New(s Services) (*server, error) {
	srv := &server{
		Services: s,
	}

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ServerAdressGRPC)
	if err != nil {
		return nil, err
	}
	srv.closer(lis.Close)
	srv.lis = lis

	// Create a gRPC server object
	srv.grpcServer = grpc.NewServer(
		grpc.ConnectionTimeout(GRPCTimeoutConnection),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	srv.closer(func() error {
		srv.grpcServer.Stop()
		return nil
	})

	// Attach Services to the server
	srv.registerServices()

	return srv, nil
}

func (s *server) registerServices() {
	api.RegisterAdminServer(s.grpcServer, s.Admin)
	api.RegisterUserServer(s.grpcServer, s.User)
	api.RegisterForumServer(s.grpcServer, s.Forum)
	api.RegisterThreadServer(s.grpcServer, s.Thread)
	api.RegisterPostServer(s.grpcServer, s.Post)
}

func (s *server) registerGatewayServices(ctx context.Context, mux *runtime.ServeMux) error {
	if err := gw_api.RegisterAdminHandlerServer(ctx, mux, s.Admin); err != nil {
		return err
	}
	if err := gw_api.RegisterUserHandlerServer(ctx, mux, s.User); err != nil {
		return err
	}
	if err := gw_api.RegisterForumHandlerServer(ctx, mux, s.Forum); err != nil {
		return err
	}
	if err := gw_api.RegisterThreadHandlerServer(ctx, mux, s.Thread); err != nil {
		return err
	}
	if err := gw_api.RegisterPostHandlerServer(ctx, mux, s.Post); err != nil {
		return err
	}

	return nil
}

func (s *server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("got signal:", <-sig)
		cancel()

		log.Println("server shutdown")
		if err := s.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Create a gRPC Gateway mux
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	// Serve the swagger-ui and swagger file
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// Register Swagger Handler
	fs := http.FileServer(http.Dir(SwaggerDir))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	// Register Gateway
	if err := s.registerGatewayServices(ctx, gwmux); err != nil {
		return err
	}

	// Create a gRPC Gateway server
	gwServer := &http.Server{
		Addr:    ServerAdressHTTP,
		Handler: mux,
	}

	s.closer(gwServer.Close)

	// Serve gRPC server
	s.group.Go(func() error {
		log.Println("start listen gRPC on", ServerAdressGRPC)
		return s.grpcServer.Serve(s.lis)
	})

	// Serve gateway server
	s.group.Go(func() error {
		log.Println("start listen HTTP on", ServerAdressHTTP)
		return gwServer.ListenAndServe()
	})

	return s.group.Wait()
}

func (s *server) closer(c closer) {
	s.closers = append(s.closers, c)
}

func (s *server) Close() (err error) {
	for i := len(s.closers) - 1; i > 0; i-- {
		err = s.closers[i]()
	}
	return err
}
