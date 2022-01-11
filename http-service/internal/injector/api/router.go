package api

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/app/router"
)

var RouteSet = wire.NewSet(router.GinRouterSet)
