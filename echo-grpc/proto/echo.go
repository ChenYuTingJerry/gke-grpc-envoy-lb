//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative echo.proto

package proto

import (
	"context"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Server for the Echo gRPC API
type Server struct {
	UnimplementedEchoServiceServer
}

// Echo the content of the request
func (s *Server) Echo(ctx context.Context, in *EchoRequest) (*EchoResponse, error) {
	log.Printf("Handling Echo request [%v] with context %v", in, ctx)
	logrus.Info("Handling Echo request [%v] with context %v", in, ctx)
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Unable to get hostname %v", err)
		hostname = ""
	}
	grpc.SendHeader(ctx, metadata.Pairs("hostname", hostname))
	return &EchoResponse{Content: in.Content}, nil
}
