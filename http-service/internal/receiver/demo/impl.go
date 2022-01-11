package demo

import (
	"net/http"

	"github.com/gin-gonic/gin"

	http2 "github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/domain/model/http"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/receiver"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/translator/demo"
)

type impl struct {
	translator demo.Translator
}

func ProvideReceiver(translator demo.Translator) Receiver {
	im := &impl{
		translator: translator,
	}
	return im
}

func (im *impl) GetRouteInfos() []receiver.RouteInfo {
	return []receiver.RouteInfo{
		{
			Method:      http.MethodGet,
			Path:        "/demo",
			Middlewares: nil,
			Handler:     im.getDemo,
		},
	}
}

func (im *impl) getDemo(c *gin.Context) {
	context, req, err := im.translator.ParseGetDemoReq(c)
	if err != nil {
		http2.ResFailResult(c, err.Code, err.ErrCode, err.Message)
		return
	}

	result, err := im.translator.Process(context, req)
	if err != nil {
		http2.ResFailResult(c, err.Code, err.ErrCode, err.Message)
		return
	}
	http2.ResSuccess(c, http.StatusOK, result)
}

func (im *impl) createDemo(c *gin.Context) {
	http2.ResSuccess(c, http.StatusCreated, map[string]string{"good": "job"})
}
