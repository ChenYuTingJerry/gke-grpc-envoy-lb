package demo

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/dispatcher/demo"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/domain/model/errs"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/translator"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/translator/demo/requests"
)

type impl struct {
	dispatcher demo.Dispatcher
}

func ProvideTranslator(dispatcher demo.Dispatcher) Translator {
	return &impl{
		dispatcher: dispatcher,
	}
}

func (im *impl) ParseGetDemoReq(c *gin.Context) (context.Context, translator.Request, *errs.AppError) {
	context := c.MustGet("ctx").(context.Context)
	var req = &requests.GetDemoReq{}
	return context, req, nil
}

func (im *impl) Process(context context.Context, req translator.Request) (interface{}, *errs.AppError) {
	switch t := req.(type) {
	case *requests.GetDemoReq:
		return im.processGetDemo(context, req.(*requests.GetDemoReq))
	default:
		return nil, errs.NewUnexpectedError(fmt.Sprintf("%v is of a type I don't know how to handle", t))
	}
}

func (im *impl) processGetDemo(context context.Context, req *requests.GetDemoReq) (interface{}, *errs.AppError) {
	result, err := im.dispatcher.GetDemo(context)
	if err != nil {
		return nil, errs.NewProcessError(err.Error())
	}
	return result, nil
}
