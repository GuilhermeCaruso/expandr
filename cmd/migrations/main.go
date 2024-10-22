package main

import (
	"flag"
	"log"

	"github.com/expandr/expandr/internal/settings"
	"github.com/expandr/expandr/migrations"
	"github.com/expandr/expandr/pkg/database"
)

func main() {
	cmd := flag.String("cmd", "", "display colorized output")
	to := flag.String("to", "", "display colorized output")

	flag.Parse()

	if *cmd != "up" && *cmd != "down" && *cmd != "init" {
		log.Fatalf("command should be one of the [up, down or init]")
		return
	}

	cfg := settings.NewConfig()

	db := database.NewDatabase(
		database.WithDSN(cfg.Db.DSN),
		database.WithMaxConns(cfg.Db.MaxConns),
		database.WithMaxIdleConns(cfg.Db.MaxIdleConns),
		database.WithMaxConnIdleLifetime(cfg.Db.MaxConnIdleLifetime),
		database.WithMaxConnLifetime(cfg.Db.MaxConnLifetime),
	)

	migrations.Migrator.Execute(db, database.MigratorParams{
		Cmd:         database.MigratorCommand(*cmd),
		VersionHash: to,
	},
	)
}
