package api

import (
	"github.com/google/wire"

	"http-service/internal/receiver"
	"http-service/internal/receiver/demo"
)

var ReceiverSet = wire.NewSet(
	demo.ProvideReceiver,
)

func ProvideReceiverList(a demo.Receiver) []receiver.Receiver {
	return []receiver.Receiver{a}
}
