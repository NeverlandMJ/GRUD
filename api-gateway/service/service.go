package service

import (
	"context"
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
	GetUserPosts(ctx context.Context, userID int) ([]response.DataResponse, error)
	GetPostByID(ctx context.Context, postID int) (response.DataResponse, error)
}

func NewService(collecterURL, grudURL string) Service {
	return Service{
		Collecter: NewGRPCClientCollect(collecterURL),
		Grud:      NewGRPCClientGrud(grudURL),
	}
}
