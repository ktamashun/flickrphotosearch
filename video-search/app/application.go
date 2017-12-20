package app

import (
	"go.uber.org/fx"
	"log"
	"os"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"context"
	"github.com/ktamashun/flickrphotosearch/video-search/app/servers"
)

type Application struct {
	*fx.App
}

func NewLogger() (*log.Logger, error) {
	logger := log.New(os.Stdout, "Log message: ", 0)
	logger.Println("Executing NewLogger")
	return logger, nil
}

func NewGrpcServer(lc fx.Lifecycle, logger *log.Logger) *grpc.Server {
	// gRPC API
	server := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(i context.Context) error {
			listener, err := net.Listen("tcp", ":28000")
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
				return err
			}

			if err := server.Serve(listener); err != nil {
				log.Fatalf("failed to serve: %v", err)
				return err
			}

			logger.Println("gRPC server is running")
			return nil
		},
		OnStop: func(i context.Context) error {

			logger.Println("Stopping gRPC server")
			server.GracefulStop()
			return nil
		},
	})

	return server
}

func RegisterReflectionService(server *grpc.Server) {
	// Register reflection service on gRPC server.
	reflection.Register(server)
}

func NewApp() *Application {
	return &Application{
		App: fx.New(
			fx.Provide(
				NewLogger,
				//NewMux,
				//NewHandler,
			),

			fx.Provide(
				NewGrpcServer,
			),

			fx.Invoke(
				//Register,
				RegisterReflectionService,
			),

			servers.ProvideVideoSearchService(),
		),
	}
}

/*
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Println("Executing NewMux")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":28080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Println("Starting Server")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Println("Stoping server")
			server.Shutdown(ctx)
			return nil
		},
	})

	return mux
}

func NewHandler(logger *log.Logger) http.Handler {
	logger.Println("Executing NewHandler")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	})
}

func Register(logger *log.Logger, mux *http.ServeMux, h http.Handler) {
	logger.Println("Executing Runner..")
	mux.Handle("/valami", h)
}
*/
