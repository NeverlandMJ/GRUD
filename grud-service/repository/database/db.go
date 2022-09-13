package database

import (
	"context"
	"github.com/NeverlandMJ/GRUD/grud-service/config"
	"github.com/NeverlandMJ/GRUD/grud-service/entity"
	"github.com/NeverlandMJ/GRUD/grud-service/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB(cfg config.Config) (*PostgresDB, error) {
	db, err := postgres.Connect(cfg)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (p *PostgresDB) GetPostsByUserID(ctx context.Context, userID int) ([]entity.Data, error) {
	query := `SELECT id, user_id, title, body FROM posts WHERE user_id=$1`

	data := make([]entity.Data, 0)
	if err := p.db.SelectContext(ctx, &data, query, userID); err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PostgresDB) GetPostByID(ctx context.Context, postID int) (entity.Data, error) {
	query := `SELECT id, user_id, title, body FROM posts WHERE id=$1`

	data := entity.Data{}
	if err := p.db.GetContext(ctx, &data, query, postID); err != nil {
		return entity.Data{}, err
	}
	return data, nil
}

func (p *PostgresDB) createPost(ctx context.Context, datas []entity.Data) error {
	query := `INSERT INTO posts (id, user_id, title, body) VALUES ($1, $2, $3, $4)`

	for _, data := range datas {
		if _, err := p.db.ExecContext(ctx, query, data.ID, data.UserID, data.Title, data.Body); err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgresDB) cleanup(ctx context.Context) error {
	query := `DELETE FROM posts`

	if _, err := p.db.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}
