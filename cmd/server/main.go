package main

import (
	"github.com/expandr/expandr/internal/server"
	"github.com/expandr/expandr/internal/settings"
	"github.com/expandr/expandr/pkg/database"
	v1 "github.com/expandr/expandr/src/v1/handlers"
)

func main() {
	cfg := settings.NewConfig()

	database.NewDatabase(
		database.WithDSN(cfg.Db.DSN),
		database.WithMaxConns(cfg.Db.MaxConns),
		database.WithMaxIdleConns(cfg.Db.MaxIdleConns),
		database.WithMaxConnIdleLifetime(cfg.Db.MaxConnIdleLifetime),
		database.WithMaxConnLifetime(cfg.Db.MaxConnLifetime),
	)

	server.NewServer(
		server.WithPort(cfg.App.Port),
	).
		RegisterVersion(1, v1.NewHandlers()).
		Listen()
}
