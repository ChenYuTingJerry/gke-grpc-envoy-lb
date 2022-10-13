package api

import (
	"github.com/google/wire"

	"http-service/internal/dispatcher/demo"
)

var DispatcherSet = wire.NewSet(
	demo.ProvideDispatcher,
)
