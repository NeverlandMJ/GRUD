package service

import (
	"context"

	"github.com/NeverlandMJ/GRUD/grud-service/entity"
	db "github.com/NeverlandMJ/GRUD/grud-service/repository/database"
)

type Service struct {
	repo db.Repository
}

func NewService(repo db.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetPostsByUserID(ctx context.Context, userID int) ([]entity.Data, error) {
	return s.repo.GetPostsByUserID(ctx, userID)
}

func (s Service) GetPostByID(ctx context.Context, postID int) (entity.Data, error) {
	return s.repo.GetPostByID(ctx, postID)
}

func (s Service) DeletePost(ctx context.Context, postID int) error {
	return s.repo.DeletePost(ctx, postID)
}

func (s Service) UpdateTitle(ctx context.Context, postID int, newTitle string) (entity.Data, error) {
	return s.repo.UpdateTitle(ctx, postID, newTitle)
}

func (s Service) UpdateBody(ctx context.Context, postID int, newBody string) (entity.Data, error) {
	return s.repo.UpdateBody(ctx, postID, newBody)
}

func (s Service) ListPosts(ctx context.Context, page, limit int) ([]entity.Data, int, error) {
	return s.repo.ListPosts(ctx, page, limit)
}
