package api

import (
	"http-service/internal/app/engine"
	"http-service/internal/app/router"
)

func InitGinEngine(r router.Router) engine.HttpEngine {
	engine := &engine.GinEngine{}
	engine.Init(r)
	return engine
}
