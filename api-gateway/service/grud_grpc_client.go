package service

import (
	"context"
	"github.com/NeverlnadMJ/GRUD/api-gateway/pkg/response"
	"github.com/NeverlnadMJ/GRUD/api-gateway/protos/grudpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type grudGRPCClient struct {
	client grudpb.GrudServiceClient
}

func NewGRPCClientGrud(url string) grudGRPCClient {
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
	client := grudpb.NewGrudServiceClient(conn)

	return grudGRPCClient{
		client: client,
	}
}

func (c grudGRPCClient) GetUserPosts(ctx context.Context, userID int) ([]response.DataResponse, error) {
	resp, err := c.client.GetPostsByUserID(ctx, &grudpb.GetUserPostsRequest{
		UserId: int32(userID),
	})

	if err != nil {
		return nil, err
	}

	data := make([]response.DataResponse, 0, len(resp.Datas))
	for _, p := range resp.Datas {
		data = append(data, response.ToDataResponse(p))
	}

	return data, nil
}

func (c grudGRPCClient) GetPostByID(ctx context.Context, postID int) (response.DataResponse, error) {
	resp, err := c.client.GetPostByID(ctx, &grudpb.GetPostByIDRequest{
		PostId: int32(postID),
	})
	if err != nil {
		return response.DataResponse{}, err
	}

	return response.ToDataResponse(resp), nil
}
