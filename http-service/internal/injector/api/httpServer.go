package api

import (
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/app/engine"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/app/router"
)

func InitGinEngine(r router.Router) engine.HttpEngine {
	engine := &engine.GinEngine{}
	engine.Init(r)
	return engine
}
