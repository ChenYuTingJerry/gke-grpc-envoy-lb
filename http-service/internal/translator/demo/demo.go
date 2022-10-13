package demo

import (
	"context"

	"github.com/gin-gonic/gin"

	"http-service/internal/lib/errs"
	"http-service/internal/translator"
)

type Translator interface {
	ParseGetDemoReq(*gin.Context) (context.Context, translator.Request, *errs.AppError)
	Process(context.Context, translator.Request) (interface{}, *errs.AppError)
}
