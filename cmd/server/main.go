package main

import (
	"github.com/expandr/expandr/internal/server"
	"github.com/expandr/expandr/internal/settings"
	v1 "github.com/expandr/expandr/internal/v1/handlers"
	"github.com/expandr/expandr/models"
	"github.com/expandr/expandr/pkg/database"
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

	database.GetDatabase().Conn.RegisterModel((*models.CollectionKey)(nil))

	server.NewServer(
		server.WithPort(cfg.App.Port),
	).
		RegisterVersion(1, v1.NewHandlers()).
		Listen()
}
