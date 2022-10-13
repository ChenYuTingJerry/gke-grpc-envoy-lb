package client

import (
	"github.com/google/wire"

	"http-service/internal/common/client"
)

var GrpcClientSet = wire.NewSet(client.NewEchoGrpcClient)
