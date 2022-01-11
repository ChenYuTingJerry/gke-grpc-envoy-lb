package demo

import "context"

type RepositoryStub struct{}

func ProvideRepositoryStub() Repository {
	return &RepositoryStub{}
}

func (r RepositoryStub) GetDemo(ctx context.Context) (interface{}, error) {
	return struct {
		Echo string
	}{
		Echo: "I am echo",
	}, nil
}
