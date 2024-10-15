package settings

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
)

type AppConfig struct {
	Port int `env:"PORT"`
}

type DbConfig struct {
	Dsn                 string `env:"DSN"`
	MaxConns            int    `env:"MAX_CONNS"`
	MaxIdleConns        int    `env:"MAX_IDLE_CONNS"`
	MaxConnLifetime     int    `env:"MAX_CONN_LIFETIME"`
	MaxConnIdleLifetime int    `env:"MAX_CONN_IDLE_LIFETIME"`
}

type Config struct {
	App AppConfig `envPrefix:"APP_"`
	Db  DbConfig  `envPrefix:"DB_"`
}

var (
	cfg  *Config
	once sync.Once
)

func NewConfig() *Config {
	once.Do(func() {
		loadedCfg, err := env.ParseAs[Config]()
		if err != nil {
			log.Fatalf("failed to load environment variables. err=%q", err)
		}
		cfg = &loadedCfg
	})

	return cfg
}

func GetConfig() *Config {
	return NewConfig()
}
