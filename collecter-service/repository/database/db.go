package database

import (
	"context"

	"github.com/NeverlandMJ/GRUD/collecter-service/config"
	"github.com/NeverlandMJ/GRUD/collecter-service/entity"
	"github.com/NeverlandMJ/GRUD/collecter-service/repository/postgres"
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

func (p *PostgresDB) SavePost(ctx context.Context, data []entity.Data) error {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	query := `INSERT INTO posts (id, user_id, title, body) VALUE ($1, $2, $3, $4)`

	for _, d := range data {
		_, err := tx.ExecContext(ctx, query, d.ID, d.UserID, d.Title, d.Body)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
