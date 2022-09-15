package database

import (
	"context"
	"github.com/NeverlandMJ/GRUD/grud-service/entity"
)

type Repository interface {
	GetPostsByUserID(ctx context.Context, userID int) ([]entity.Data, error)
	GetPostByID(ctx context.Context, postID int) (entity.Data, error)
	DeletePost(ctx context.Context, postID int) error
	UpdateTitle(ctx context.Context, postID int, newTitle string) (entity.Data, error)
	UpdateBody(ctx context.Context, postID int, newBody string) (entity.Data, error)
	ListPosts(ctx context.Context, page, limit int) ([]entity.Data, int, error)
}
