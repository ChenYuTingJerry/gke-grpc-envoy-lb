package domain

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/domain/repository/demo"
)

var RepositorySet = wire.NewSet(
	demo.ProvideRepositoryGrpc,
)
