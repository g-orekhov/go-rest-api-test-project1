package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Postgres *pgxpool.Pool
}

func (db *DB) ConnectDB() error {
	//TODO: make configurations
	var err error
	db.Postgres, err = pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return err
	}
	return nil
}
