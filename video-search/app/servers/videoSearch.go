package servers

import (
	"log"
	"google.golang.org/grpc"
	"github.com/ktamashun/flickrphotosearch/video-search/apis"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"go.uber.org/fx"
)

func ProvideVideoSearchService() fx.Option {
	return fx.Options(
		fx.Provide(NewVideoServer),
		fx.Invoke(RegisterVideoSearchServer),
	)
}

func RegisterVideoSearchServer(logger *log.Logger, server *grpc.Server, videoServer *VideoServer) {
	logger.Println("Registering video search server")
	apis.RegisterVideoServer(server, videoServer)
}

type VideoServer struct {
	logger *log.Logger
}

func NewVideoServer(logger *log.Logger) *VideoServer {
	return &VideoServer{
		logger: logger,
	}
}

func (server *VideoServer) Search(ctx context.Context, request *apis.VideoSearchRequest) (*apis.VideoSearchResponse, error) {
	server.logger.Println("Hurray!!!")


	return nil, status.Errorf(codes.Unimplemented, "Video Search service is not yet implemented, sorry.. :(")
}
