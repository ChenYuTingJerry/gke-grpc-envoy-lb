package engine

import (
	"net/http"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/app/router"
)

type HttpEngine interface {
	Init(r router.Router)
	StartServer() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
