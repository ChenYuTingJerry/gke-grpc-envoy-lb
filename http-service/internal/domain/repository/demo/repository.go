package demo

import "context"

type Repository interface {
	GetDemo(ctx context.Context) (interface{}, error)
}
