package service

import (
	"context"
	"github.com/NeverlnadMJ/GRUD/api-gateway/pkg/request"
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

func (c grudGRPCClient) GetUserPosts(ctx context.Context, userID int) ([]response.Post, error) {
	resp, err := c.client.GetPostsByUserID(ctx, &grudpb.GetUserPostsRequest{
		UserID: int32(userID),
	})

	if err != nil {
		return nil, err
	}

	data := make([]response.Post, 0, len(resp.Posts))
	for _, p := range resp.Posts {
		data = append(data, response.ToDataResponse(p))
	}

	return data, nil
}

func (c grudGRPCClient) GetPostByID(ctx context.Context, postID int) (response.Post, error) {
	resp, err := c.client.GetPostByID(ctx, &grudpb.GetPostByIDRequest{
		PostID: int32(postID),
	})
	if err != nil {
		return response.Post{}, err
	}

	return response.ToDataResponse(resp), nil
}

func (c grudGRPCClient) DeletePost(ctx context.Context, postID int) error {
	_, err := c.client.DeletePost(ctx, &grudpb.DeletePostRequest{
		PostID: int32(postID),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c grudGRPCClient) UpdateTitle(ctx context.Context, req request.UpdatePostRequest) (response.Post, error) {
	resp, err := c.client.UpdateTitle(ctx, &grudpb.UpdateTitleRequest{
		PostID:   int32(req.PostID),
		NewTitle: req.New,
	})
	if err != nil {
		return response.Post{}, err
	}

	return response.ToDataResponse(resp), nil
}

func (c grudGRPCClient) UpdateBody(ctx context.Context, req request.UpdatePostRequest) (response.Post, error) {
	resp, err := c.client.UpdateBody(ctx, &grudpb.UpdateBodyRequest{
		PostID:  int32(req.PostID),
		NewBody: req.New,
	})
	if err != nil {
		return response.Post{}, err
	}
	return response.ToDataResponse(resp), nil
}

func (c grudGRPCClient) ListPosts(ctx context.Context, page, limit int) ([]response.Post, int, error) {
	resp, err := c.client.ListPosts(ctx, &grudpb.ListPostsRequest{
		Page:  int32(page),
		Limit: int32(limit),
	})
	if err != nil {
		return nil, 0, err
	}

	data := make([]response.Post, 0, len(resp.Data.Posts))
	for _, p := range resp.Data.Posts {
		data = append(data, response.ToDataResponse(p))
	}

	return data, int(resp.All), nil
}
