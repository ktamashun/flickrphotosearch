package main

import (
	"log"
	"net/http"

	"github.com/ktamashun/flickrphotosearch/backend/app"
	"github.com/ktamashun/flickrphotosearch/backend/apis"
	"github.com/ktamashun/flickrphotosearch/backend/app/servers"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"fmt"
)

const (
	restPort = ":8080"
	grpcPort = ":18000"
)

func main() {
	// JSONAPI
	go func() {
		router := app.CreateRouter()
		log.Fatal(http.ListenAndServe(restPort, router))
	}()

	fmt.Println("hello grpc world")
	// gRPC API
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	apis.RegisterPhotoServer(s, &servers.PhotoServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
