package engine

import (
	"net/http"

	"http-service/internal/app/router"
)

type HttpEngine interface {
	Init(r router.Router)
	StartServer() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
