package database

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once     sync.Once
	database *Database
)

type Database struct {
	Conn *gorm.DB
}

func NewDatabase(opts ...Option) *Database {
	once.Do(func() {

		internalConfig := newConfig(opts...)
		dialect := postgres.Open(internalConfig.dsn)
		db, err := gorm.Open(dialect, &gorm.Config{})

		if err != nil {
			panic(err)
		}

		sqldb, err := db.DB()

		if err := sqldb.Ping(); err != nil {
			log.Fatalf("failed to connect to the database: %v", err)
		}

		sqldb.SetMaxOpenConns(internalConfig.maxConns)
		sqldb.SetMaxIdleConns(internalConfig.maxIdleConns)
		sqldb.SetConnMaxIdleTime(internalConfig.maxConnIdleLifetime)
		sqldb.SetConnMaxLifetime(internalConfig.maxConnLifetime)

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
