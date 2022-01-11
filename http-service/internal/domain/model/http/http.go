package http

import (
	"github.com/gin-gonic/gin"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/domain/model/errs"
)

type FailResult struct {
	Code errs.ErrCode `json:"errCode"`
	Msg  string       `json:"errMessage"`
}

func ResSuccess(c *gin.Context, status int, v interface{}) {
	ResJSON(c, status, v)
}

func ResFailResult(c *gin.Context, status int, errCode errs.ErrCode, msg string) {
	ResJSON(c, status, FailResult{
		Code: errCode,
		Msg:  msg,
	})
}

func ResJSON(c *gin.Context, status int, v interface{}) {
	c.Writer.Header().Add("Content-Type", "application/json")
	c.AbortWithStatusJSON(status, v)
}
