package database

import (
	"database/sql"
	"log"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	once     sync.Once
	database *Database
)

type Database struct {
	Conn *bun.DB
}

func NewDatabase(opts ...Option) *Database {
	once.Do(func() {

		internalConfig := newConfig(opts...)

		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(internalConfig.dsn)))

		sqldb.SetMaxOpenConns(internalConfig.maxConns)
		sqldb.SetMaxIdleConns(internalConfig.maxIdleConns)
		sqldb.SetConnMaxLifetime(internalConfig.maxConnLifetime)
		sqldb.SetConnMaxIdleTime(internalConfig.maxConnIdleLifetime)

		db := bun.NewDB(sqldb, pgdialect.New())

		if err := db.Ping(); err != nil {
			log.Fatalf("failed to connect to the database: %v", err)
		}

		database = &Database{
			Conn: db,
		}
	})

	return database
}

func GetDatabase() *Database {
	if database == nil {
		log.Fatal("database not initialized. Call NewDatabase first")
	}

	return database
}
