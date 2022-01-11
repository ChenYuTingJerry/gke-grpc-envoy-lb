package injector

import (
	log "github.com/sirupsen/logrus"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/configs"
)

func Run() {
	// prevent init run before testing https://github.com/golang/go/issues/31859
	configs.InitConfigs()
	// init injector
	injector, cleanup, err := BuildInjector()
	if err != nil {
		panic(err.Error())
	}
	// start config watcher

	// start http server
	engine := injector.HttpEngine
	err = engine.StartServer()
	if err != nil {
		cleanup()
		log.WithField("err", err).Panic("Fail to start watchers")
	}
}
