package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func ConfigureDb() (*sql.DB, error) {
	dsn := "postgres://user:pass@localhost:5432/library?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
