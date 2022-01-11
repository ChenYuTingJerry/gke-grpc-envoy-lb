package injector

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/app/engine"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	HttpEngine engine.HttpEngine
}
