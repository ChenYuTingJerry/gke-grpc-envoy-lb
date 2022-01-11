package requests

import (
	"context"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/domain/model/errs"
)

type GetDemoReq struct{}

func (g *GetDemoReq) Validate(context context.Context) *errs.AppError {
	return nil
}

func (g *GetDemoReq) Extract(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
