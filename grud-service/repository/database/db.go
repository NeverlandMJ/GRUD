package database

import (
	"github.com/NeverlandMJ/GRUD/grud-service/config"
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



