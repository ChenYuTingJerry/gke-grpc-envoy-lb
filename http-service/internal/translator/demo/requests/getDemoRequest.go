package requests

import (
	"context"

	"http-service/internal/lib/errs"
)

type GetDemoReq struct{}

func (g *GetDemoReq) Validate(context context.Context) *errs.AppError {
	return nil
}

func (g *GetDemoReq) Extract(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
