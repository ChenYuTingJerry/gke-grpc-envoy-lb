package api

import (
	"github.com/google/wire"

	"http-service/internal/translator/demo"
)

var TranslatorSet = wire.NewSet(
	demo.ProvideTranslator,
)
