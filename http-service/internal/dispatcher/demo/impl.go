package demo

import (
	"context"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/domain/repository/demo"
)

type impl struct {
	demoRepo demo.Repository
}

func ProvideDispatcher(demoRepo demo.Repository) Dispatcher {
	return &impl{
		demoRepo: demoRepo,
	}
}

func (i *impl) GetDemo(ctx context.Context) (interface{}, error) {
	return i.demoRepo.GetDemo(ctx)
}
