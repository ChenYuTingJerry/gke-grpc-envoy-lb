//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"

	"http-service/internal/injector/api"
	"http-service/internal/injector/client"
	"http-service/internal/injector/domain"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// api
		api.InitGinEngine,
		api.RouteSet,
		api.ReceiverSet,
		api.ProvideReceiverList,
		api.TranslatorSet,
		api.DispatcherSet,

		// Domain
		domain.RepositorySet,

		// client
		client.GrpcClientSet,

		//Injector
		InjectorSet,
	)
	return new(Injector), nil, nil
}
