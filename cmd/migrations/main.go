package main

import (
	"github.com/expandr/expandr/internal/settings"
	"github.com/expandr/expandr/migrations"
	"github.com/expandr/expandr/pkg/database"
)

func main() {
	cfg := settings.NewConfig()

	db := database.NewDatabase(
		database.WithDSN(cfg.Db.DSN),
		database.WithMaxConns(cfg.Db.MaxConns),
		database.WithMaxIdleConns(cfg.Db.MaxIdleConns),
		database.WithMaxConnIdleLifetime(cfg.Db.MaxConnIdleLifetime),
		database.WithMaxConnLifetime(cfg.Db.MaxConnLifetime),
	)

	migrations.Migrator.Execute(db, database.UP)
}
