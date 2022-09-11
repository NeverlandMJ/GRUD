package service

import "context"

type Service struct {
	Collecter
}

type Collecter interface {
	SavePosts(ctx context.Context, pages int) error
}

func NewService(collecterURL string) Service {
	return Service{
		Collecter: NewGRPCClientCollect(collecterURL),
	}
}
