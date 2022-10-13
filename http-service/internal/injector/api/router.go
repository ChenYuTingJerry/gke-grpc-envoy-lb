package api

import (
	"github.com/google/wire"

	"http-service/internal/app/router"
)

var RouteSet = wire.NewSet(router.GinRouterSet)
