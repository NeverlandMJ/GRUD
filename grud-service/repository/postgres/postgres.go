package postgres

import (
	"fmt"

	"github.com/NeverlandMJ/GRUD/grud-service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB,
		),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
