package api

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/dispatcher/demo"
)

var DispatcherSet = wire.NewSet(
	demo.ProvideDispatcher,
)
