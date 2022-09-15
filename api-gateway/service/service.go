package service

import (
	"context"
	"github.com/NeverlnadMJ/GRUD/api-gateway/pkg/request"
	"github.com/NeverlnadMJ/GRUD/api-gateway/pkg/response"
)

type Service struct {
	Collecter
	Grud
}

type Collecter interface {
	SavePosts(ctx context.Context, pages int) error
}

type Grud interface {
	GetUserPosts(ctx context.Context, userID int) ([]response.Post, error)
	GetPostByID(ctx context.Context, postID int) (response.Post, error)
	DeletePost(ctx context.Context, postID int) error
	UpdateTitle(ctx context.Context, req request.UpdatePostRequest) (response.Post, error)
	UpdateBody(ctx context.Context, req request.UpdatePostRequest) (response.Post, error)
	ListPosts(ctx context.Context, page, limit int) ([]response.Post, int, error)
}

func NewService(collecterURL, grudURL string) Service {
	return Service{
		Collecter: NewGRPCClientCollect(collecterURL),
		Grud:      NewGRPCClientGrud(grudURL),
	}
}
