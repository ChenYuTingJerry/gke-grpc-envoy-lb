package demo

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"http-service/internal/common/server"
	"http-service/internal/receiver"
	"http-service/internal/translator/demo"
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
		server.ResFailResult(c, err.Code, err.ErrCode, err.Message)
		return
	}

	result, err := im.translator.Process(context, req)
	if err != nil {
		server.ResFailResult(c, err.Code, err.ErrCode, err.Message)
		return
	}
	server.ResSuccess(c, http.StatusOK, result)
}

func (im *impl) createDemo(c *gin.Context) {
	server.ResSuccess(c, http.StatusCreated, map[string]string{"good": "job"})
}
