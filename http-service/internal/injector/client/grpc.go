package client

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/client/grpc"
)

var GrpcClientSet = wire.NewSet(grpc.NewEchoGrpcClient)
