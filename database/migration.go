package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sqlx.DB, path string) error {
	goose.SetVerbose(false)
	goose.SetTableName("db_versions")

	return goose.Up(db.DB, path)
}
