package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/NeverlandMJ/GRUD/grud-service/config"
	"github.com/NeverlandMJ/GRUD/grud-service/entity"
	"github.com/NeverlandMJ/GRUD/grud-service/errs"
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

func (p *PostgresDB) ListPosts(ctx context.Context, page, limit int) ([]entity.Data, int, error) {
	count, err := p.count(ctx)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT * FROM posts OFFSET $1 LIMIT $2`
	offset := (page - 1) * limit
	data := make([]entity.Data, 0)

	if err = p.db.SelectContext(ctx, &data, query, offset, limit); err != nil {
		return nil, 0, err
	}

	return data, count, nil
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

		if errors.Is(err, sql.ErrNoRows) {
			return entity.Data{}, errs.ErrPostNotFound
		}
		return entity.Data{}, err
	}

	return data, nil
}

func (p *PostgresDB) DeletePost(ctx context.Context, postID int) error {
	query := `DELETE FROM posts WHERE id = $1`

	_, err := p.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) UpdateTitle(ctx context.Context, postID int, newTitle string) (entity.Data, error) {
	query := `UPDATE posts SET title=$1 WHERE id=$2 RETURNING id, user_id, title, body`

	updated := entity.Data{}

	err := p.db.QueryRowContext(ctx, query, newTitle, postID).Scan(&updated.ID, &updated.UserID, &updated.Title, &updated.Body)
	if err != nil {
		return entity.Data{}, err
	}

	return updated, nil
}

func (p *PostgresDB) UpdateBody(ctx context.Context, postID int, newBody string) (entity.Data, error) {
	query := `UPDATE posts SET body=$1 WHERE id=$2 RETURNING id, user_id, title, body`

	updated := entity.Data{}

	err := p.db.QueryRowContext(ctx, query, newBody, postID).Scan(&updated.ID, &updated.UserID, &updated.Title, &updated.Body)
	if err != nil {
		return entity.Data{}, err
	}

	return updated, nil
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

func (p *PostgresDB) count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM posts`

	var count int
	if err := p.db.GetContext(ctx, &count, query); err != nil {
		return 0, err
	}

	return count, nil
}

func (p *PostgresDB) cleanup(ctx context.Context) error {
	query := `DELETE FROM posts`

	if _, err := p.db.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}
