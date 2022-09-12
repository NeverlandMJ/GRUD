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

	stmt, err := tx.Prepare("INSERT INTO posts (id, user_id, title, body) VALUES ($1, $2, $3, $4)")
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, d := range data {
		_, err = stmt.ExecContext(ctx, d.ID, d.UserID, d.Title, d.Body)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	stmt.Close()
	return nil
}
