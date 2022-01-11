package translator

import (
	"context"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/domain/model/errs"
)

type Request interface {
	//Validate is for checking request parameters format
	Validate(ctx context.Context) *errs.AppError
	//Extract is for extracting some request parameters into specific values.
	Extract(ctx context.Context)
}

type Response interface {
	//Convert is for converting business model to response model
	Convert(ctx context.Context, input interface{}) *errs.AppError
}
