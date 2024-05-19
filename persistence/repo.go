package persistence

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Repository struct {
	db *sql.DB
}

func NewRepo(dbURL string) (*Repository, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	err = makeMigrations(db)
	if err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("making migration: %w", err)
	}
	return &Repository{
		db: db,
	}, nil
}

func makeMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://schema",
		"postgres", driver)
	if err != nil {
		return err
	}
	return m.Up()

}
