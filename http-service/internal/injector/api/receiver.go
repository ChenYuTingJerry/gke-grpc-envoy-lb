package api

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/receiver"
	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/receiver/demo"
)

var ReceiverSet = wire.NewSet(
	demo.ProvideReceiver,
)

func ProvideReceiverList(a demo.Receiver) []receiver.Receiver {
	return []receiver.Receiver{a}
}
