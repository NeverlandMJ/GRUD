package service

import (
	"context"

	db "github.com/NeverlandMJ/GRUD/collecter-service/repository/database"
	"github.com/NeverlandMJ/GRUD/collecter-service/entity"
)

type Service struct {
	repo db.Repository
}

func NewService(repo db.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) SavePost(ctx context.Context, data []entity.Data) error {
	return s.repo.SavePost(ctx, data)
}
