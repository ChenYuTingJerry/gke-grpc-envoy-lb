package api

import (
	"github.com/google/wire"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/translator/demo"
)

var TranslatorSet = wire.NewSet(
	demo.ProvideTranslator,
)
