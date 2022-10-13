package demo

import (
	"context"

	api "http-service/internal/common/client/proto"
)

type RepositoryGrpc struct {
	client api.EchoClient
}

func ProvideRepositoryGrpc(client api.EchoClient) Repository {
	return &RepositoryGrpc{
		client: client,
	}
}

func (r RepositoryGrpc) GetDemo(ctx context.Context) (interface{}, error) {
	req := &api.EchoRequest{
		Content: "Aloha",
	}
	res, err := r.client.Echo(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
