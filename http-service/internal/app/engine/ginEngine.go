package engine

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"http-service/configs"
	"http-service/internal/app/router"
)

type GinEngine struct {
	engine *gin.Engine
}

func (g *GinEngine) Init(r router.Router) {
	app := gin.Default()
	r.RegisterAPI(app)
	if configs.C.App.EnableProfile {
		fmt.Println("profiling")
		// TODO if we want to always enable profile need to deal with whitelist for this api
		// default is "debug/pprof"
		pprof.Register(app, "debug/pprof")
	}
	g.engine = app
}

func (g *GinEngine) StartServer() error {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", configs.C.App.Port),
		Handler:      g.engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (g *GinEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	g.engine.ServeHTTP(w, req)
}
