package main

import (
	"log"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/echo-grpc/health"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/echo-grpc/proto"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterEchoServiceServer(grpcServer, &proto.Server{})
	grpc_health_v1.RegisterHealthServer(grpcServer, &health.Server{})
	logrus.Infof("Listening for Echo on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
