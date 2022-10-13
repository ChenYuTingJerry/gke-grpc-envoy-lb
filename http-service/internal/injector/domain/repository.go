package domain

import (
	"github.com/google/wire"

	"http-service/internal/domain/repository/demo"
)

var RepositorySet = wire.NewSet(
	demo.ProvideRepositoryGrpc,
)
