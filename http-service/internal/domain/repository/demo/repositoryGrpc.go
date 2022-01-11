package demo

import (
	"context"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/echo-grpc/proto"
)

type RepositoryGrpc struct {
	client proto.EchoServiceClient
}

func ProvideRepositoryGrpc(client proto.EchoServiceClient) Repository {
	return &RepositoryGrpc{
		client: client,
	}
}

func (r RepositoryGrpc) GetDemo(ctx context.Context) (interface{}, error) {
	req := &proto.EchoRequest{
		Content: "Aloha",
	}
	res, err := r.client.Echo(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
