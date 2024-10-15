package database

import "time"

type config struct {
	maxConns            int
	maxIdleConns        int
	maxConnLifetime     time.Duration
	maxConnIdleLifetime time.Duration
}

type Option func(*config)

func WithMaxConns(maxConns int) Option {
	return func(c *config) {
		c.maxConns = maxConns
	}
}

func WithMaxIdleConns(maxIdleConns int) Option {
	return func(c *config) {
		c.maxIdleConns = maxIdleConns
	}
}

func WithMaxConnLifetime(minutes int) Option {
	return func(c *config) {
		c.maxConnLifetime = time.Duration(minutes) * time.Minute
	}
}

func WithMaxConnIdleLifetime(minutes int) Option {
	return func(c *config) {
		c.maxConnIdleLifetime = time.Duration(minutes) * time.Minute
	}
}

func (c *config) defaults() {
	c.maxConns = 25
	c.maxIdleConns = 5
	c.maxConnIdleLifetime = 5 * time.Minute
	c.maxConnLifetime = 15 * time.Minute
}

func newConfig(opts ...Option) *config {
	c := new(config)

	c.defaults()

	for _, opt := range opts {
		opt(c)
	}

	return c
}
