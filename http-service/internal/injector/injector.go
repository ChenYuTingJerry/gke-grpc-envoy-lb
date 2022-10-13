package injector

import (
	"github.com/google/wire"

	"http-service/internal/app/engine"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	HttpEngine engine.HttpEngine
}
