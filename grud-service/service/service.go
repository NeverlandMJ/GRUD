package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/NeverlandMJ/GRUD/grud-service/entity"
	"github.com/NeverlandMJ/GRUD/grud-service/errs"
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
	d, err := s.repo.GetPostByID(ctx, postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Data{}, errs.ErrPostNotFound
		}
		return entity.Data{}, err
	}

	return d, nil
}
