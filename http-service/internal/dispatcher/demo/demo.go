package demo

import "context"

type Dispatcher interface {
	GetDemo(ctx context.Context) (interface{}, error)
}
