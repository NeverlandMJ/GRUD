package service

import (
	"context"
	"time"

	"github.com/NeverlnadMJ/GRUD/api-gateway/v1/collectpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type collecterGRPCClient struct {
	client collectpb.CollecterServiceClient
}

func NewGRPCClientCollect(url string) collecterGRPCClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	client := collectpb.NewCollecterServiceClient(conn)

	return collecterGRPCClient{
		client: client,
	}
}

func (c collecterGRPCClient) SavePosts(ctx context.Context, pages int) error {
	_, err := c.client.SavePosts(ctx, &collectpb.DownloadReq{
		Page: int32(pages),
	})
	if err != nil {
		return err
	}

	return nil
}
