package receiver

import (
	"github.com/gin-gonic/gin"

	"http-service/internal/common/middlewares/common"
)

type Receiver interface {
	GetRouteInfos() []RouteInfo
}

type RouteInfo struct {
	Method      string
	Path        string
	Middlewares []gin.HandlerFunc
	Handler     gin.HandlerFunc
}

func (r *RouteInfo) GetFlow() []gin.HandlerFunc {
	var flow []gin.HandlerFunc
	// Append common translations here
	flow = append(flow, common.AddContext)

	// Append handler
	flow = append(flow, r.Handler)
	return flow
}
