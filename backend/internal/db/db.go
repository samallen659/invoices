package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(connectionURI string) (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", connectionURI)
	if err != nil {
		return nil, fmt.Errorf("failed to open Postgres connection: %w", err)
	}
	return conn, nil
}
