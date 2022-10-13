package translator

import (
	"context"

	"http-service/internal/lib/errs"
)

type Request interface {
	//Validate is for checking request parameters format
	Validate(ctx context.Context) *errs.AppError
	//Extract is for extracting some request parameters into specific values.
	Extract(ctx context.Context)
}

type Response interface {
	//Convert is for converting business model to response model
	Convert(ctx context.Context, input interface{}) *errs.AppError
}
