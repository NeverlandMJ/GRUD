package database

import (
	"context"

	"github.com/NeverlandMJ/GRUD/collecter-service/entity"
)

type Repository interface {
	SavePost(ctx context.Context, data []entity.Data) error
}
