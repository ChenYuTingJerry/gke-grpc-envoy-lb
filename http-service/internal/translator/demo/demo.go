package demo

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/domain/model/errs"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/translator"
)

type Translator interface {
	ParseGetDemoReq(*gin.Context) (context.Context, translator.Request, *errs.AppError)
	Process(context.Context, translator.Request) (interface{}, *errs.AppError)
}
