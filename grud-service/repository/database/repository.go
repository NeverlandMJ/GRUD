package database

import (
	"context"
	"github.com/NeverlandMJ/GRUD/grud-service/entity"
)

type Repository interface {
	GetPostsByUserID(ctx context.Context, userID int) ([]entity.Data, error)
	GetPostByID(ctx context.Context, postID int) (entity.Data, error)
}
